package repository

import (
	"cat-social-be/helper"
	"cat-social-be/model/domain"
	requestdto "cat-social-be/model/dto/request"
	responsedto "cat-social-be/model/dto/response"
	"database/sql"
	"log"
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
	query := "SELECT name, password_hash, email FROM users WHERE email = $1"
	resultUser := domain.User{}
	errorQuery := tx.QueryRow(query, user.Email).Scan(&resultUser.Name, &resultUser.Password, &resultUser.Email)
	if errorQuery != nil {
		log.Fatal(errorQuery)
	}

	err = VerifyPassword(user.Password, resultUser.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User login failed",
			"data":    "Password is incorrect",
		})
		return
	}

	token, _ := helper.GenerateToken(user.Email)
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
	query := "INSERT INTO users (email, name, password_hash) VALUES ($1,$2,$3) RETURNING email, name"
	resultUser := domain.User{}
	errorQuery := tx.QueryRow(query, user.Email, user.Name, user.Password).Scan(&resultUser.Email, &resultUser.Name)
	if errorQuery != nil {
		log.Fatal(errorQuery)
	}
	token, _ := helper.GenerateToken(user.Email)
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
	query := "SELECT email FROM users WHERE email = $1"
	resultUser := domain.User{}
	tx.QueryRow(query, email).Scan(&resultUser.Email)
	return resultUser.Email != ""
}

func FindIdByEmail(c *gin.Context, tx *sql.DB, email string) int {
	query := "SELECT id FROM users WHERE email = $1"
	resultUser := domain.User{}
	tx.QueryRow(query, email).Scan(&resultUser.Id)
	return resultUser.Id
}
