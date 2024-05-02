package controller

import (
	"cat-social-be/helper"
	requestdto "cat-social-be/model/dto/request"
	repository "cat-social-be/repository/user"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	userCreateRequest := requestdto.UserCreateRequest{}
	c.ShouldBindJSON(&userCreateRequest)
	// helper.ReadFromRequestBody(request, &userCreateRequest)

	loginResponse, _ := repository.Login(c, db, userCreateRequest)
	token, _ := helper.GenerateToken("USER_CAT")
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Berhasil login user", "data": loginResponse, "token": token})
	// resultResponse := responsedto.DefaultResponse{
	// 	Message: "OK",
	// 	Data:    categoryResponse,
	// }

	// helper.WriteToResponseBody(writer, resultResponse)
}

func Register(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	userCreateRequest := requestdto.UserCreateRequest{}
	c.ShouldBindJSON(&userCreateRequest)
	userToken, _ := helper.ExtractTokenRole(c)
	fmt.Println("userToken", userToken)
	if userToken != "USER_CAT" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "You are unauthorized to access this resource, this resource for USER_CAT user",
		})
		return
	}
	loginResponse, _ := repository.Register(c, db, userCreateRequest)
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Berhasil register user", "data": loginResponse})
}
