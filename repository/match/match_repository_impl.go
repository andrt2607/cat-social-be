package repository

import (
	"cat-social-be/helper"
	requestdto "cat-social-be/model/dto/request"
	responsedto "cat-social-be/model/dto/response"
	userRepository "cat-social-be/repository/user"
	"cat-social-be/model/domain"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetMatches(c *gin.Context, tx *sql.DB) (responsedto.DefaultResponse, error) {
	loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	idUser := userRepository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	queryTes := "select l.id, l.message, l.created_at, u.name, u.email, u.created_at, c.id, c.name, c.race, c.sex, c.description, c.age_in_month, c.image_urls, c.is_matched, c.created_at, c2.id , c2.name , c2.race , c2.sex , c2.description , c2.age_in_month , c2.image_urls , c2.is_matched , c2.created_at from likes l join users u on l.owner_id = u.id join cats c ON l.liked_cat_id = c.id join cats c2 on l.cat_id = c2.id where l.owner_id = $1 or l.liked_owner_id = $1 order by l.created_at desc"
	rows, errorQuery := tx.Query(queryTes, idUser)
	if errorQuery != nil {
		log.Fatal(errorQuery)
	}
	var listData []responsedto.MatchGetResponse
	for rows.Next() {
		var data responsedto.MatchGetResponse
		rows.Scan(&data.ID, &data.Message, &data.CreatedAt, &data.IssuedBy.Name, &data.IssuedBy.Email, &data.IssuedBy.CreatedAt, &data.MatchCatDetail.ID, &data.MatchCatDetail.Name, &data.MatchCatDetail.Race, &data.MatchCatDetail.Sex, &data.MatchCatDetail.Description, &data.MatchCatDetail.AgeInMonth, &data.MatchCatDetail.ImageUrls, &data.MatchCatDetail.HasMatched, &data.MatchCatDetail.CreatedAt, &data.UserCatDetail.ID, &data.UserCatDetail.Name, &data.UserCatDetail.Race, &data.UserCatDetail.Sex, &data.UserCatDetail.Description, &data.UserCatDetail.AgeInMonth, &data.UserCatDetail.ImageUrls, &data.UserCatDetail.HasMatched, &data.UserCatDetail.CreatedAt)
		fmt.Print("data", data)
		listData = append(listData, data)
	}
	response := responsedto.DefaultResponse{
		Message: "success",
		Data:    listData,
	}
	c.JSON(http.StatusOK, response)
	return response, nil
}

func ValidateCreateMatch(c *gin.Context, tx *sql.DB, req requestdto.MatchCreateRequest) (domain.Cat, domain.Cat, string, int, string, error) {
	loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))

	query := "SELECT id, name, owner_id, sex, is_matched, is_deleted FROM cats WHERE id in ($1, $2)"
	fmt.Println(query)
	rows, err := tx.Query(query, req.UserCatId, req.MatchCatId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	userCat := domain.Cat{}
	matchCat := domain.Cat{}

	// var checks []domain.cat
	for rows.Next() {
		check := domain.Cat{}
		err := rows.Scan(
			&check.Id, 
			&check.Name, 
			&check.OwnerId, 
			&check.Sex, 
			&check.IsMatched, 
			&check.IsDeleted,
		)
		helper.PanicIfError(err)
		if strconv.Itoa(check.Id) == req.UserCatId {
			if check.OwnerId != idUser{
				log.Fatal(err)
				err_message := fmt.Sprintf("cat id %s is not belong to the user %s", check.Id, loggedUserEmail)
				return userCat, matchCat, "", http.StatusBadRequest, err_message, nil
			} else{
				userCat = check
			}
		} else {
			matchCat = check
		}
	}
	if userCat.Sex == matchCat.Sex {
		log.Fatal(err)
		err_message := fmt.Sprintf("your cat id %s gender %s is the same with match cat id %s gender %s", userCat.Id, userCat.Sex, matchCat.Id, matchCat.Sex)
		return userCat, matchCat, "", http.StatusBadRequest, err_message, nil
	}
	if (userCat.IsMatched == true) || (matchCat.IsMatched == true) {
		log.Fatal(err)
		err_message := fmt.Sprintf("neither cat id %s and %s has been matched", userCat.Id, matchCat.Id)

		return userCat, matchCat, "", http.StatusBadRequest, err_message, nil
	}
	if userCat.OwnerId == matchCat.OwnerId {
		log.Fatal(err)
		err_message := fmt.Sprintf("cat id %s and %s is from the same owner", userCat.Id, matchCat.Id)
		return userCat, matchCat, "", http.StatusBadRequest, err_message, nil
	}

	return userCat, matchCat, req.Message, 0, "", nil
}

func CreateMatch(c *gin.Context, tx *sql.DB, catUser domain.Cat, matchUser domain.Cat, matchMessage string) (responsedto.DefaultResponse, error) {
	//get id user from email token jwt
	// loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	// idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	query := "INSERT INTO likes (owner_id, cat_id, liked_owner_id, liked_cat_id, is_approved, message) VALUES ($1, $2, $3, $4, NULL, $5) RETURNING id, created_at"

	resultMatch := domain.Match{}
	//run query insert
	err := tx.QueryRow(
		query, 
		catUser.OwnerId, 
		catUser.Id, 
		matchUser.OwnerId, 
		matchUser.Id, 
		matchMessage,
	).Scan(&resultMatch.Id, &resultMatch.CreatedAt)
	//handle error
	if err != nil {
		log.Fatal(err)
	}
	// defer rows.Close()
	response := responsedto.DefaultResponse{
		Message: "success",
		Data: responsedto.MatchCreateResponse{
			Id:        resultMatch.Id,
			CreatedAt: resultMatch.CreatedAt,
		},
	}
	return response, nil
}

func ApproveMatch(c *gin.Context, tx *sql.DB, req requestdto.MatchApproveRequest) (responsedto.DefaultResponse, error) {
	//get id user from email token jwt
	// loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	// idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	query_update := "UPDATE likes SET is_approved = 'APPROVED', updated_at = $1 WHERE id = $2 RETURNING id, cat_id, liked_cat_id, created_at, updated_at"
	resultMatch := domain.Match{}

	//run query update
	err := tx.QueryRow(
		query_update, 
		time.Now(), 
		req.MatchId,
	).Scan(&resultMatch.Id, &resultMatch.CatId, &resultMatch.LikedCatId, &resultMatch.CreatedAt, &resultMatch.UpdatedAt)
	
	//handle error
	if err != nil {
		log.Fatal(err)
	}

	query_delete := "DELETE FROM likes WHERE (cat_id IN ($1, $2) or liked_cat_id IN ($1, $2)) and id <> $3"
	
	//run query delete
	err_delete := tx.QueryRow(query_delete, resultMatch.CatId, resultMatch.LikedCatId, resultMatch.Id)
	if err_delete != nil {
		log.Fatal(err_delete)
	}

	query_update := "UPDATE likes SET is_matched = True WHERE id IN ($1, $2)"
	err_update := tx.QueryRow(
		query_update, 
		resultMatch.CatId,
		resultMatch.LikedCatId,
	)
	if err_update != nil {
		log.Fatal(err_update)
	}
	// defer rows.Close()

	response := responsedto.DefaultResponse{
		Message: "success",
		Data: responsedto.MatchApproveResponse{
			Id:        		resultMatch.Id,
			CatId:        	resultMatch.CatId,
			LikedCatId:     resultMatch.LikedCatId,
			CreatedAt: 		resultMatch.CreatedAt,
			UpdatedAt:		resultMatch.UpdatedAt,
		},
	}
	return response, nil
}
