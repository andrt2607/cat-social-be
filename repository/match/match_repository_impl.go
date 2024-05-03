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
func CreateMatch(c *gin.Context, tx *sql.DB, req map[string]interface{}) (responsedto.DefaultResponse, error) {
	//get id user from email token jwt
	loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	// idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	query := "INSERT INTO likes (owner_email, cat_id, liked_owner_email, liked_cat_id, is_approved, message) VALUES ($1, $2, $3, $4, NULL, $5) RETURNING id, created_at"

	resultMatch := domain.Match{}
	//run query insert
	err := tx.ExecContext(
		c,
		query, 
		req["user_cat"].
		OwnerEmail, 
		req["user_cat"].Id, 
		req["user_cat"].OwnerEmail, 
		req["user_cat"].Id, 
		req["message"]
	).Scan(&resultMatch.Id, &resultMatch.CreatedAt)
	//handle error
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	response := responsedto.DefaultResponse{
		Message: "success",
		Data: responsedto.MatchCreateResponse{
			Id:        resultMatch.Id,
			CreatedAt: resultMatch.CreatedAt,
		},
	}
	return response, nil
}

func ValidateCreateMatch(c *gin.Context, tx *sql.DB, req requestdto.MatchCreateRequest) (domain.Cat, error) {
	loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	// idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))

	query := "SELECT id, name, owner_email, sex, is_matched, is_deleted FROM cats WHERE cat_id in ($1, $2)"
	rows, err := tx.QueryContext(
		ctx, 
		SQL, 
		req.userCatId, 
		req.matchCatId
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	checks := map[string]interface{}

	// var checks []domain.cat
	for rows.Next() {
		check := domain.cat{}
		err := rows.Scan(
			&check.Id, 
			&check.Name, 
			&check.OwnerEmail, 
			&check.Sex, 
			&check.isMatched, 
			&check.isDeleted
		)
		helper.PanicIfError(err)
		if check.Id == req.userCatId{
			if check.OwnerEmail != loggedUserEmail{
				log.Fatal(err)
				err_message := fmt.Sprintf("cat id %s is not belong to the user %s", check.Id, loggedUserEmail)
				return http.StatusBadRequest, err_message
			} else{
				checks["user_cat"] = check
			}
		} else {
			checks["match_cat"] = check
		}
	}
	if checks["user_cat"].Sex == checks["match_cat"].Sex {
		log.Fatal(err)
		err_message := fmt.Sprintf("your cat id %s gender %s is the same with match cat id %s gender %s", checks["user_cat"].Id, checks["user_cat"].Sex, checks["match_cat"].Id, checks["match_cat"].Sex)
		return nil, http.StatusBadRequest ,err_message
	}
	if (checks["user_cat"].isMatched == True) || (checks["match_cat"].isMatched == True) {
		log.Fatal(err)
		err_message := fmt.Sprintf("neither cat id %s and %s has been matched", checks["user_cat"].Id, checks["match_cat"].Id)

		return nil, http.StatusBadRequest ,err_message
	}
	if checks["user_cat"].OwnerEmail == checks["match_cat"].OwnerEmail {
		log.Fatal(err)
		err_message := fmt.Sprintf("cat id %s and %s is from the same owner", checks["user_cat"].Id, checks["match_cat"].Id)
		return nil, http.StatusBadRequest ,err_message
	}
	checks["message"] = req.message

	return checks, nil, nil

func ApproveMatch(c *gin.Context, tx *sql.DB, req requestdto.MatchApproveRequest) (responsedto.DefaultResponse, error) {
	//get id user from email token jwt
	loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	// idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	query_update := "UPDATE likes SET is_approved = 'APPROVED', updated_at = $1 WHERE id = $2 RETURNING id, cat_id, liked_cat_id, created_at, updated_at"
	resultMatch := domain.Match{}

	//run query update
	_, err := tx.ExecContext(
		c, 
		query_update, 
		time.Now(), 
		req.matchId
	).Scan(
		&resultMatch.Id, 
		&resultMatch.catId, 
		&resultMatch.likedCatId,
		&resultMatch.CreatedAt, 
		&resultMatch.UpdatedAt, 
	)
	//handle error
	if err != nil {
		log.Fatal(err)
	}
	query_delete = "DELETE FROM likes WHERE (cat_id IN ($1, $2) or liked_cat_id IN ($1, $2) and id <> $3"
	//run query delete
	_, err := tx.Exec(
		c, 
		query_delete, 
		resultMatch.catId,
		resultMatch.likedCatId,
		resultMatch.Id
	)
	query_updated_cat = "UPDATE likes SET is_matched = True WHERE id IN ($1, $2)"
	_, err := tx.Exec(
		c, 
		query_updated_cat, 
		resultMatch.catId,
		resultMatch.likedCatId
	)
	defer rows.Close()

	response := responsedto.DefaultResponse{
		Message: "success",
		Data: responsedto.MatchApproveResponse{
			Id:        		resultMatch.Id,
			CatId:        	resultMatch.catId,
			LikedCatId:     resultMatch.likedCatId,
			CreatedAt: 		resultMatch.CreatedAt,
			UpdatedAt:		resultMatch.UpdatedAt
		},
	}
	return response, nil
}
