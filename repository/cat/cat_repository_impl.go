package repository

import (
	"cat-social-be/helper"
	"cat-social-be/model/domain"
	requestdto "cat-social-be/model/dto/request"
	responsedto "cat-social-be/model/dto/response"
	repository "cat-social-be/repository/user"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateCat(c *gin.Context, tx *sql.DB, user requestdto.CatCreateRequest) {
	//get id user from email token jwt
	loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	query := "INSERT INTO cats (name, race, sex, age_in_month, description, image_urls ,owner_id, is_matched, is_deleted) VALUES ($1, $2, $3, $4, $5,  $6 , $7, $8, $9) RETURNING id, created_at"
	fmt.Println("ini query insertnya ", query)
	fmt.Println("ini request nya ", user)
	var arrayText string
	values := user.ImageUrls
	arrayText += "{"
	for i, value := range values {
		arrayText += fmt.Sprintf(`"%s"`, value)
		if i < len(values)-1 {
			arrayText += ","
		}
	}
	arrayText += "}"
	resultCat := domain.Cat{}
	fmt.Println("ini array text nya ", arrayText)
	//run query insert
	err := tx.QueryRow(query, user.Name, user.Race, user.Sex, user.AgeInMonth,
		user.Description, arrayText, idUser, false, false).Scan(&resultCat.Id, &resultCat.CreatedAt)
	//handle error
	if err != nil {
		// log.Fatal(err)
		fmt.Println("errornya masuk sini : ", err)
		response := responsedto.DefaultResponse{
			Message: "failed to create cat",
			Data:    errors.New(err.Error()),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// convertId := strconv.Itoa(resultCat.Id)
	// catResponse :=
	// catResponse.Id = resultCat.Id
	// catResponse.CreatedAt = resultCat.CreatedAt.Format(time.RFC3339)
	response := responsedto.DefaultResponse{
		Message: "success",
		Data: responsedto.CatCreateResponse{
			Id:        resultCat.Id,
			CreatedAt: resultCat.CreatedAt.Format(time.RFC3339),
		},
	}
	c.JSON(http.StatusCreated, response)
	fmt.Println("response : ", response)
	// return response, nil
}

func GetCats(c *gin.Context, tx *sql.DB) {
	loggedUserEmail, _ := helper.ExtractTokenEmail(c)
	idUser := repository.FindIdByEmail(c, tx, loggedUserEmail.(string))
	query := "SELECT id, name, race, sex, age_in_month, description, image_urls, is_matched, created_at FROM cats"
	filtersString := []string{}
	//add filter id
	filterId := c.Query("id")
	if filterId != "" {
		filtersString = append(filtersString, fmt.Sprintf(" id = %s", filterId))
	}
	//add fillter race
	filterRace := c.Query("race")
	if filterRace != "" {
		filtersString = append(filtersString, fmt.Sprintf(" race = '%s'", filterRace))
	}
	//add fillter sex
	filterSex := c.Query("sex")
	if filterSex != "" {
		filtersString = append(filtersString, fmt.Sprintf(" sex = '%s'", filterSex))
	}
	//add fillter by search / name
	filterSearchName := c.Query("search")
	if filterSearchName != "" {
		filtersString = append(filtersString, " name LIKE '%"+filterSearchName+"%'")
	}
	//add fillter hasMatched
	filterHasMatched := c.Query("hasMatched") == "true"
	// fmt.Print("filterHasMatched : ", filterHasMatched)
	if filterHasMatched {
		filtersString = append(filtersString, fmt.Sprintf(" is_matched = %s", strconv.FormatBool(filterHasMatched)))
	} else {
		filtersString = append(filtersString, fmt.Sprintf(" is_matched = %s", strconv.FormatBool(filterHasMatched)))
	}
	//add fillter ageInMonth
	filterAgeInMonth := c.Query("ageInMonth")
	fmt.Println("filterAgeInMonth : ", filterAgeInMonth)
	if filterAgeInMonth != "" {
		_, err := strconv.Atoi(filterAgeInMonth)
		numberOnlyAgeInMonth := filterAgeInMonth[1:]
		if err != nil && filterAgeInMonth[0] != '=' {
			filtersString = append(filtersString, fmt.Sprintf(" age_in_month %s%s", string(filterAgeInMonth[0]), numberOnlyAgeInMonth))
		} else {
			filtersString = append(filtersString, fmt.Sprintf(" age_in_month = %s", numberOnlyAgeInMonth))
		}
	}
	//add fillter owned
	filterOwned := c.Query("owned") == "true"
	if filterOwned {
		filtersString = append(filtersString, fmt.Sprintf(" owner_id = %d", idUser))
	}
	//concat all filter
	filterGetString := ""
	if len(filtersString) > 0 {
		filterGetString = " WHERE " + strings.Join(filtersString, " AND ") + " AND is_deleted = false"
	} else {
		filterGetString = " WHERE is_deleted = false"

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
	fmt.Println("query get filter nya : ", queryWithFilter)
	if errorQuery != nil {
		fmt.Print("error query : ", errorQuery)
		return
		// log.Fatal(errorQuery)
	}
	var listCat []responsedto.CatGetResponse
	for rows.Next() {
		// var textArray string
		var imageUrlsStr string
		var cat responsedto.CatGetResponse
		var createdAtBefore time.Time
		var idBefore int
		err := rows.Scan(&idBefore, &cat.Name, &cat.Race,
			&cat.Sex, &cat.AgeInMonth, &cat.Description,
			&imageUrlsStr, &cat.HasMatched, &createdAtBefore)
		cat.Id = strconv.Itoa(idBefore)
		cat.CreatedAt = createdAtBefore.Format(time.RFC3339)
		imageUrls := string(imageUrlsStr)
		// Parsing nilai array teks
		imageUrls = strings.Trim(imageUrls, "{}") // Menghapus kurung kurawal
		imageUrlsArr := strings.Split(imageUrls, ",")
		cat.ImageUrls = imageUrlsArr
		fmt.Println("imageUrlsArr : ", cat.ImageUrls)
		if err != nil {
			fmt.Print("error scan : ", err)
			// log.Fatal(err)
			return
		}
		// cat.ImageUrls = strings.Split(textArray[1:len(textArray)-1], ",")
		// fmt.Print("cat imageurls : ", cat.ImageUrls)
		listCat = append(listCat, cat)
	}
	if len(listCat) == 0 && filterId != "" {
		response := responsedto.DefaultResponse{
			Message: "Cat is already deleted or not found",
			Data:    []string{},
		}
		c.JSON(http.StatusOK, response)
		return
	}
	fmt.Println("sampai sini end")
	response := responsedto.DefaultResponse{
		Message: "success",
		Data:    listCat,
	}
	c.JSON(http.StatusOK, response)
	// return response, nil
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
	var arrayText string
	values := user.ImageUrls
	arrayText += "{"
	for i, value := range values {
		arrayText += fmt.Sprintf(`"%s"`, value)
		if i < len(values)-1 {
			arrayText += ","
		}
	}
	arrayText += "}"
	//query update
	query := "UPDATE cats SET name = $1, race = $2, sex = $3, age_in_month = $4, description = $5, image_urls = $6, owner_id = $7 , updated_at = $9 WHERE id = $8 RETURNING id, created_at"
	resultCat := domain.Cat{}
	err := tx.QueryRow(query, user.Name, user.Race, user.Sex, user.AgeInMonth,
		user.Description, arrayText, idUser, idCat, time.Now()).Scan(&resultCat.Id, &resultCat.CreatedAt)
	//handle error
	if err != nil {
		fmt.Println("errornya masuk sini update cat : \n", err)
		log.Fatal(err)
	}
	response := responsedto.DefaultResponse{
		Message: "successfully update cat",
		Data: responsedto.CatCreateResponse{
			Id:        resultCat.Id,
			CreatedAt: resultCat.CreatedAt.Format(time.RFC3339),
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
		fmt.Println("errornya masuk sini check is already matched : \n", err)
		log.Fatal(err)
	}
	if resultCat.IsMatched && inputSex != resultCat.Sex {
		return true
	}
	return false
}

func DeleteCat(c *gin.Context, tx *sql.DB, user requestdto.CatCreateRequest) (responsedto.DefaultResponse, error) {
	fmt.Println("masuk sini delete cat")
	// query := "DELETE FROM cats WHERE id = $1"
	query := "UPDATE cats SET is_deleted = true WHERE id = $1"
	_, err := tx.Exec(query, c.Param("id"))
	if err != nil {
		fmt.Println("errornya masuk sini delete cat : \n", err)
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
		fmt.Println("errornya masuk sini find cat by id : \n", err)
		log.Fatal(err)
	}
	return resultCat, nil
}
