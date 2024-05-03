package controller

import (
	"cat-social-be/helper"
	requestdto "cat-social-be/model/dto/request"
	catRepository "cat-social-be/repository/cat"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleInternalServerError(c *gin.Context, err error) {
	fmt.Println("Internal Server Error:", err)
	// Mengirim respons dengan status HTTP 500 (Internal Server Error)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal Server Error",
	})
}

func GetCats(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	//call repository
	catRepository.GetCats(c, db)
}

func CreateCat(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	catCreateRequest := requestdto.CatCreateRequest{}
	c.ShouldBindJSON(&catCreateRequest)
	//validasi input
	if err := helper.ValidateStruct(&catCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//call repository
	catCreateResponse, _ := catRepository.CreateCat(c, db, catCreateRequest)
	c.JSON(http.StatusCreated, catCreateResponse)
}

func UpdateCat(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	catCreateRequest := requestdto.CatCreateRequest{}
	c.ShouldBindJSON(&catCreateRequest)
	//validasi input
	if err := helper.ValidateStruct(&catCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//validasi id cat
	_, err := catRepository.FindCatById(c, db)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Id cat is not found",
		})
		return

	}
	//call repository
	catCreateResponse, _ := catRepository.UpdateCat(c, db, catCreateRequest)
	c.JSON(http.StatusCreated, catCreateResponse)
}

func DeleteCat(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	catCreateRequest := requestdto.CatCreateRequest{}
	c.ShouldBindJSON(&catCreateRequest)
	//validasi id cat
	_, err := catRepository.FindCatById(c, db)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Id cat is not found",
		})
		return

	}
	//call repository
	catCreateResponse, _ := catRepository.DeleteCat(c, db, catCreateRequest)
	c.JSON(http.StatusCreated, catCreateResponse)
}
