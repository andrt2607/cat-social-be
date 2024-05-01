package repository

import (
	"cat-social-be/helper"
	"cat-social-be/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	fmt.Print(user.Email, user.Password)
	// SQL := "select email, password from user where email = 'aliefazuka123@gmail.com'"
	SQL := "SELECT * FROM user WHERE name = 'alip'"
	rows, err := tx.QueryContext(ctx, SQL, user.Email)
	helper.PanicIfError(err)
	defer rows.Close()

	resultUser := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Email, &user.Password)
		helper.PanicIfError(err)
		fmt.Println(resultUser)
		return resultUser, nil
	} else {
		return resultUser, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "update user set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Id)
	helper.PanicIfError(err)

	return user, nil
}
