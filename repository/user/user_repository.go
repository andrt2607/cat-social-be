package repository

import (
	"cat-social-be/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	Login(ctx context.Context, tx *sql.Tx, User domain.User) (domain.User, error)
	Register(ctx context.Context, tx *sql.Tx, User domain.User) (domain.User, error)
}
