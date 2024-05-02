package repository

import (
	"cat-social-be/helper"
	"cat-social-be/model/domain"
	requestdto "cat-social-be/model/dto/request"
	responsedto "cat-social-be/model/dto/response"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Login(c *gin.Context, tx *sql.DB, user requestdto.UserLoginRequest) {
	var err error
	query := "SELECT * FROM users WHERE email = $1"
	resultUser := domain.User{}
	tx.QueryRow(query, user.Email).Scan(&resultUser.Id, &resultUser.Email, &resultUser.Name, &resultUser.Password)

	err = VerifyPassword(user.Password, resultUser.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User login failed",
			"data":    "Password is incorrect",
		})
		return
	}

	token, _ := helper.GenerateToken("USER_CAT")
	response := responsedto.DefaultResponse{
		Message: "User logged successfully",
		Data: responsedto.AuthResponse{
			Email:       resultUser.Email,
			Name:        resultUser.Name,
			AccessToken: token,
		},
	}
	c.JSON(http.StatusOK, response)
}

func Register(c *gin.Context, tx *sql.DB, user requestdto.UserCreateRequest) (responsedto.DefaultResponse, error) {
	query := "INSERT INTO users (email, name, password) VALUES ($1,$2,$3) RETURNING id"
	resultUser := domain.User{}
	tx.QueryRow(query, user.Email, user.Name, user.Password).Scan(&resultUser.Id, &resultUser.Email, &resultUser.Name, &resultUser.Password)
	token, _ := helper.GenerateToken("USER_CAT")
	response := responsedto.DefaultResponse{
		Message: "User registered successfully",
		Data: responsedto.AuthResponse{
			Email:       user.Email,
			Name:        user.Name,
			AccessToken: token,
		},
	}
	return response, nil
}

func IsEmailExist(c *gin.Context, tx *sql.DB, email string) bool {
	query := "SELECT * FROM users WHERE email = $1"
	resultUser := domain.User{}
	tx.QueryRow(query, email).Scan(&resultUser.Id, &resultUser.Email, &resultUser.Name, &resultUser.Password)
	return resultUser.Email != ""
}
