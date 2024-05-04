package controller

import (
	"cat-social-be/helper"
	requestdto "cat-social-be/model/dto/request"
	matchRepository "cat-social-be/repository/match"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleInternalServerError(c *gin.Context, err error) {
	// Mengirim respons dengan status HTTP 500 (Internal Server Error)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal Server Error",
	})
}

func CreateMatch(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	matchCreateRequest := requestdto.MatchCreateRequest{}
	c.ShouldBindJSON(&matchCreateRequest)
	//validasi input
	if err := helper.ValidateStruct(&matchCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validasi Match Request
	catUser, matchUser, matchMessage, err_code, err_message, err_validate := matchRepository.ValidateCreateMatch(c, db, matchCreateRequest)
	if err_validate != nil {
		c.JSON(err_code, gin.H{
			"error": err_message,
		})
		return
	}
	//call repository
	matchCreateResponse, _ := matchRepository.CreateMatch(c, db, catUser, matchUser, matchMessage)
	c.JSON(http.StatusCreated, matchCreateResponse)
}

func GetMatches(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	matchRepository.GetMatches(c, db)
}

func ApproveMatch(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	matchApproveRequest := requestdto.MatchApproveRequest{}
	c.ShouldBindJSON(&matchApproveRequest)
	//validasi input
	if err := helper.ValidateStruct(&matchApproveRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//call repository
	responseMessage, responseCode := matchRepository.ApproveMatch(c, db, matchApproveRequest)
	c.JSON(responseCode, responseMessage)
}

func DeleteMatch(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	matchRepository.DeleteMatch(c, db)
}

func RejectMatch(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	matchApproveRequest := requestdto.MatchApproveRequest{}
	c.ShouldBindJSON(&matchApproveRequest)
	//validasi input
	if err := helper.ValidateStruct(&matchApproveRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//call repository
	matchApproveResponse, statusCode, _ := matchRepository.RejectMatch(c, db, matchApproveRequest)
	c.JSON(statusCode, matchApproveResponse)
}
