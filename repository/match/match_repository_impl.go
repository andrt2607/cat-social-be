package repository

import (
	"cat-social-be/helper"
	"cat-social-be/model/domain"
	requestdto "cat-social-be/model/dto/request"
	responsedto "cat-social-be/model/dto/response"
	userRepository "cat-social-be/repository/user"
	"database/sql"
	"errors"
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
	idUser := userRepository.FindIdByEmail(c, tx, loggedUserEmail.(string))

	query := "SELECT id, name, owner_id, sex, is_matched, is_deleted FROM cats WHERE id in ($1, $2)"
	fmt.Println("ini query validate", query)
	rows, err := tx.Query(query, req.UserCatId, req.MatchCatId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	userCat := domain.Cat{}
	matchCat := domain.Cat{}

	fmt.Println("userCat kosong", userCat)
	fmt.Println("matchCat kosong", matchCat)

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
		fmt.Println("isi check", check)
		helper.PanicIfError(err)
		if strconv.Itoa(check.Id) == req.UserCatId {
			if check.OwnerId != strconv.Itoa(idUser) {
				log.Fatal(err)
				err_message := fmt.Sprintf("cat id %s is not belong to the user %s", check.Id, loggedUserEmail)
				return userCat, matchCat, "", http.StatusBadRequest, err_message, errors.New(err_message)
			} else {
				userCat = check
			}
		} else {
			matchCat = check
		}
	}
	fmt.Println("userCat : ", userCat)
	if (userCat.IsMatched) || (matchCat.IsMatched) {
		// log.Fatal(err)
		err_message := fmt.Sprintf("neither cat id %v and %v has been matched", userCat.Id, matchCat.Id)

		return userCat, matchCat, "", http.StatusBadRequest, err_message, errors.New(err_message)
	}
	if userCat.OwnerId == matchCat.OwnerId {
		// log.Fatal(err)
		err_message := fmt.Sprintf("cat id %v and %v is from the same owner", userCat.Id, matchCat.Id)
		return userCat, matchCat, "", http.StatusBadRequest, err_message, errors.New(err_message)
	}
	if userCat.Sex == matchCat.Sex {
		// log.Fatal(err)
		err_message := fmt.Sprintf("your cat id %v gender %v is the same with match cat id %v gender %v", userCat.Id, userCat.Sex, matchCat.Id, matchCat.Sex)
		return userCat, matchCat, "", http.StatusBadRequest, err_message, errors.New(err_message)
	}
	fmt.Println("message method validate : ", req.Message)
	return userCat, matchCat, req.Message, 0, "", nil
}

func CreateMatch(c *gin.Context, tx *sql.DB, catUser domain.Cat, matchUser domain.Cat, matchMessage string) (responsedto.DefaultResponse, error) {
	//get id user from email token jwt
	// loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	// idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	query := "INSERT INTO likes (owner_id, cat_id, liked_owner_id, liked_cat_id, approval_status, message) VALUES ($1, $2, $3, $4, 'pending', $5) RETURNING id, created_at"
	fmt.Println("ini query create match", query)
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
	fmt.Println("resultMatch", resultMatch)
	// defer rows.Close()
	response := responsedto.DefaultResponse{
		Message: "success",
		Data: responsedto.MatchCreateResponse{
			Id:        resultMatch.Id,
			CreatedAt: resultMatch.CreatedAt,
		},
	}
	fmt.Println("response create match", response)
	return response, nil
}

func ApproveMatch(c *gin.Context, tx *sql.DB, req requestdto.MatchApproveRequest) (responsedto.DefaultResponse, error) {
	//get id user from email token jwt
	// loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	// idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	query_update := "UPDATE likes SET approval_status = 'approved', updated_at = $1 WHERE id = $2 RETURNING id, cat_id, liked_cat_id, created_at, updated_at"
	resultMatch := domain.Match{}

	//run query update
	err := tx.QueryRow(
		query_update,
		time.Now(),
		req.MatchId,
	).Scan(&resultMatch.Id, &resultMatch.CatId, &resultMatch.LikedCatId, &resultMatch.CreatedAt, &resultMatch.UpdatedAt)
	fmt.Println("query update ", query_update)
	//handle error
	if err != nil {
		log.Fatal(err)
	}

	query_delete := "DELETE FROM likes WHERE ((cat_id = $1 or cat_id = $2) or (liked_cat_id = $1 or liked_cat_id = $2)) and id <> $3"

	//run query delete
	_, err_delete := tx.Exec(query_delete, resultMatch.CatId, resultMatch.LikedCatId, resultMatch.Id)
	if err_delete != nil {
		log.Fatal(err_delete)
	}
	fmt.Println("query delete ", query_delete)

	fmt.Println("query update cat value : ", resultMatch.CatId, resultMatch.LikedCatId)
	query_update_cat := "UPDATE cats SET is_matched = true WHERE id = $1 or id = $2"
	_, err_update := tx.Exec(query_update_cat, resultMatch.CatId, resultMatch.LikedCatId)
	fmt.Println("query update cat ", query_update_cat)
	if err_update != nil {
		log.Fatal(err_update)
	}
	// defer rows.Close()

	response := responsedto.DefaultResponse{
		Message: "success",
		Data: responsedto.MatchApproveResponse{
			Id:         resultMatch.Id,
			CatId:      resultMatch.CatId,
			LikedCatId: resultMatch.LikedCatId,
			CreatedAt:  resultMatch.CreatedAt,
			UpdatedAt:  resultMatch.UpdatedAt,
		},
	}
	return response, nil
}

func DeleteMatch(c *gin.Context, tx *sql.DB) {
	//check match matchId is already approved / reject
	queryCheckMatchIdExist := "SELECT id FROM likes WHERE id = $1"
	var id int
	errCheckMatchIdExist := tx.QueryRow(queryCheckMatchIdExist, c.Param("id")).Scan(&id)
	if errCheckMatchIdExist != nil {
		err_message := "matchId is not found"
		response := responsedto.DefaultResponse{
			Message: err_message,
			Data:    nil,
		}
		c.JSON(http.StatusNotFound, response)
	}
	//check match matchId is already approved / reject
	queryCheckIsAlreadyApproved := "SELECT approval_status FROM likes WHERE id = $1"
	var approvalStatus string
	errCheckIsAlreadyApproved := tx.QueryRow(queryCheckIsAlreadyApproved, c.Param("id")).Scan(&approvalStatus)
	if errCheckIsAlreadyApproved != nil {
		log.Fatal(errCheckIsAlreadyApproved)
	}
	if approvalStatus == "approved" || approvalStatus == "rejected" {
		response := responsedto.DefaultResponse{
			Message: "failed delete match",
			Data:    "matchId is already approved / reject",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	idUser := userRepository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	queryGetIssuerId := "SELECT owner_id FROM likes WHERE id = $1"
	var ownerId int
	errGetIssuer := tx.QueryRow(queryGetIssuerId, c.Param("id")).Scan(&ownerId)
	if errGetIssuer != nil {
		log.Fatal(errGetIssuer)
	}
	// validasi logged user not owner id cat
	if idUser != ownerId {
		err_message := "you are not authorized to delete this match"
		response := responsedto.DefaultResponse{
			Message: err_message,
			Data:    nil,
		}
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	//query delete match
	query := "DELETE FROM likes WHERE id = $1"
	_, errQueryDelete := tx.Exec(query, c.Param("id"))
	if errQueryDelete != nil {
		log.Fatal(errQueryDelete)
	}
	response := responsedto.DefaultResponse{
		Message: "success delete match",
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}
