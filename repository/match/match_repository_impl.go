package repository

import (
	"cat-social-be/helper"
	responsedto "cat-social-be/model/dto/response"
	userRepository "cat-social-be/repository/user"
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
