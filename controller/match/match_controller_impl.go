package controller

import (
	"cat-social-be/helper"
	requestdto "cat-social-be/model/dto/request"
	repository "cat-social-be/repository/match"
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

func CreateMatch(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	userLoginRequest := requestdto.UserLoginRequest{}
	c.ShouldBindJSON(&userLoginRequest)
	if err := helper.ValidateStruct(&userLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if !repository.IsEmailExist(c, db, userLoginRequest.Email) {
	// 	c.JSON(http.StatusConflict, gin.H{
	// 		"error": "User Not Found",
	// 	})
	// 	return
	// }

	repository.GetMatches(c, db)
}

//	func GetMatches(c *gin.Context) {
//		defer func() {
//			if err := recover(); err != nil {
//				handleInternalServerError(c, fmt.Errorf("%v", err))
//			}
//		}()
//		db := c.MustGet("db").(*sql.DB)
//		userCreateRequest := requestdto.UserCreateRequest{}
//		c.ShouldBindJSON(&userCreateRequest)
//		if err := helper.ValidateStruct(&userCreateRequest); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		if repository.IsEmailExist(c, db, userCreateRequest.Email) {
//			c.JSON(http.StatusConflict, gin.H{
//				"error": "Email already exist",
//			})
//			return
//		}
//		hashedPassword, err := helper.HashPassword(userCreateRequest.Password)
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		userCreateRequest.Password = hashedPassword
//		registerResponse, _ := repository.Register(c, db, userCreateRequest)
//		c.JSON(http.StatusCreated, registerResponse)
//	}
func GetMatches(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
	repository.GetMatches(c, db)
}

func ApproveMatch(c *gin.Context) {

}

func RejectMatch(c *gin.Context) {

}

func DeleteMatch(c *gin.Context) {

}
