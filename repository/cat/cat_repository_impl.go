package repository

import (
	"cat-social-be/helper"
	"cat-social-be/model/domain"
	requestdto "cat-social-be/model/dto/request"
	responsedto "cat-social-be/model/dto/response"
	repository "cat-social-be/repository/user"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateCat(c *gin.Context, tx *sql.DB, user requestdto.CatCreateRequest) (responsedto.DefaultResponse, error) {
	//get id user from email token jwt
	loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	query := "INSERT INTO cats (name, race, sex, age_in_month, description, image_urls ,owner_id, is_matched, is_deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, created_at"
	resultCat := domain.Cat{}
	//run query insert
	err := tx.QueryRow(query, user.Name, user.Race, user.Sex, user.AgeInMonth,
		user.Description, user.ImageUrls, idUser, false, false).Scan(&resultCat.Id, &resultCat.CreatedAt)
	//handle error
	if err != nil {
		log.Fatal(err)
	}
	response := responsedto.DefaultResponse{
		Message: "success",
		Data: responsedto.CatCreateResponse{
			Id:        resultCat.Id,
			CreatedAt: resultCat.CreatedAt,
		},
	}
	return response, nil
}

func GetCats(c *gin.Context, tx *sql.DB) (responsedto.DefaultResponse, error) {
	query := "SELECT id, name, race, sex, age_in_month, description, image_urls, is_matched, created_at FROM cats WHERE is_deleted = false"
	filtersString := []string{}
	//add filter id
	filterId := c.Query("id")
	if filterId != "" {
		filtersString = append(filtersString, fmt.Sprintf("id = %s", filterId))
	}
	//add fillter race
	filterRace := c.Query("race")
	if filterRace != "" {
		filtersString = append(filtersString, fmt.Sprintf("race = '%s'", filterRace))
	}
	//add fillter sex
	filterSex := c.Query("sex")
	if filterSex != "" {
		filtersString = append(filtersString, fmt.Sprintf("sex = '%s'", filterSex))
	}
	//add fillter hasMatched
	filterHasMatched := c.Query("hasMatched")
	if filterSex != "" {
		filtersString = append(filtersString, fmt.Sprintf("is_matched = %s", filterHasMatched))
	}
	//add fillter ageInMonth
	filterAgeInMonth := c.Query("ageInMonth")
	if filterAgeInMonth != "" {
		filtersString = append(filtersString, fmt.Sprintf("age_in_month = %s", filterAgeInMonth))
	}
	//add fillter owned
	filterOwned := c.Query("owned")
	if filterOwned != "" {
		filtersString = append(filtersString, fmt.Sprintf("owner_id = %s", filterOwned))
	}
	//concat all filter
	filterGetString := ""
	if len(filtersString) > 0 {
		filterGetString = strings.Join(filtersString, " AND ")
	}
	//add filter order by created_at desc
	filterGetString += " ORDER BY created_at DESC"
	//filter limit offset
	filterLimit := c.Query("limit")
	filterOffset := c.Query("offset")
	if filterLimit != "" && filterOffset != "" {
		filterGetString += fmt.Sprintf(" LIMIT %s OFFSET %s", filterLimit, filterOffset)
	} else {
		filterGetString += " LIMIT 5 OFFSET 0"
	}
	//concat query with filter
	queryWithFilter := query + filterGetString
	//run query select
	rows, errorQuery := tx.Query(queryWithFilter)
	if errorQuery != nil {
		log.Fatal(errorQuery)
	}
	var listCat []responsedto.CatGetResponse
	for rows.Next() {
		var cat responsedto.CatGetResponse
		rows.Scan(&cat.Id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.ImageUrls, &cat.HasMatched, &cat.CreatedAt)
		listCat = append(listCat, cat)
	}
	response := responsedto.DefaultResponse{
		Message: "success",
		Data:    listCat,
	}
	c.JSON(http.StatusOK, response)
	return response, nil
}

func UpdateCat(c *gin.Context, tx *sql.DB, user requestdto.CatCreateRequest) (responsedto.DefaultResponse, error) {
	//get id user from email token jwt
	idCat := c.Param("id")
	//check is already matched
	if checkIsAlreadyMatched(tx, idCat, string(user.Sex)) {
		response := responsedto.DefaultResponse{
			Message: "failed",
			Data:    "sex is edited when cat is already requested to match",
		}
		return response, nil
	}
	//get id user from email token jwt
	loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	//query update
	query := "UPDATE cats SET name = $1, race = $2, sex = $3, age_in_month = $4, description = $5, image_urls = $6, owner_id = $7 , updated_at = $9 WHERE id = $8 RETURNING id, created_at"
	resultCat := domain.Cat{}
	err := tx.QueryRow(query, user.Name, user.Race, user.Sex, user.AgeInMonth,
		user.Description, user.ImageUrls, idUser, idCat, time.Now()).Scan(&resultCat.Id, &resultCat.CreatedAt)
	//handle error
	if err != nil {
		log.Fatal(err)
	}
	response := responsedto.DefaultResponse{
		Message: "successfully update cat",
		Data: responsedto.CatCreateResponse{
			Id:        resultCat.Id,
			CreatedAt: resultCat.CreatedAt,
		},
	}
	return response, nil
}

func checkIsAlreadyMatched(tx *sql.DB, idCat string, inputSex string) bool {
	query := "SELECT sex, is_matched FROM cats WHERE id = $1"
	resultCat := domain.Cat{}
	err := tx.QueryRow(query, idCat).Scan(&resultCat.Sex, &resultCat.IsMatched)
	//handle error
	if err != nil {
		log.Fatal(err)
	}
	if resultCat.IsMatched && inputSex != resultCat.Sex {
		return true
	}
	return false
}

func DeleteCat(c *gin.Context, tx *sql.DB, user requestdto.CatCreateRequest) (responsedto.DefaultResponse, error) {
	// query := "DELETE FROM cats WHERE id = $1"
	query := "UPDATE cats SET is_deleted = true WHERE id = $1"
	_, err := tx.Exec(query, c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	response := responsedto.DefaultResponse{
		Message: "success",
		Data:    nil,
	}
	return response, nil
}

func FindCatById(c *gin.Context, tx *sql.DB) (domain.Cat, error) {
	query := "SELECT id FROM cats WHERE id = $1"
	resultCat := domain.Cat{}
	err := tx.QueryRow(query, c.Param("id")).Scan(&resultCat.Id)
	if err != nil {
		log.Fatal(err)
	}
	return resultCat, nil
}
