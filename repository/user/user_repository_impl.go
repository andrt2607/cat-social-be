package repository

import (
	"cat-social-be/model/domain"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "SELECT * FROM cobaz WHERE email = $1"
	tx.QueryRow(query, user.Email).Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	resultUser := domain.User{}
	resultUser.Id = user.Id
	resultUser.Email = user.Email
	resultUser.Name = user.Name
	resultUser.Password = user.Password
	return resultUser, nil
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	var pk int
	query := `INSERT INTO cobaz (email, name, password) VALUES ($1,$2,$3) RETURNING id`
	tx.QueryRow(query, user.Email, user.Name, user.Password).Scan(&pk)
	resultUser := domain.User{}
	resultUser.Id = pk
	resultUser.Email = user.Email
	resultUser.Name = user.Name
	resultUser.Password = user.Password
	return resultUser, nil
}
