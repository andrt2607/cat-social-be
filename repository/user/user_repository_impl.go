package repository

import (
	"cat-social-be/model/domain"
	requestdto "cat-social-be/model/dto/request"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Login(c *gin.Context, tx *sql.DB, user requestdto.UserCreateRequest) (domain.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	resultUser := domain.User{}
	tx.QueryRow(query, user.Email).Scan(&resultUser.Id, &resultUser.Email, &resultUser.Name, &resultUser.Password)
	fmt.Println(resultUser)

	return resultUser, nil
}

func Register(c *gin.Context, tx *sql.DB, user requestdto.UserCreateRequest) (domain.User, error) {
	query := "INSERT INTO users (email, name, password) VALUES ($1,$2,$3) RETURNING id"
	resultUser := domain.User{}
	tx.QueryRow(query, user.Email, user.Name, user.Password).Scan(&resultUser.Id, &resultUser.Email, &resultUser.Name, &resultUser.Password)
	fmt.Println("done tambah", resultUser)
	return resultUser, nil
}
