package service

import (
	"cat-social-be/helper"
	"cat-social-be/model/domain"
	requestdto "cat-social-be/model/dto/request"
	responsedto "cat-social-be/model/dto/response"
	repository "cat-social-be/repository/user"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(UserRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Login(ctx context.Context, request requestdto.UserCreateRequest) responsedto.DefaultResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	user, _ = service.UserRepository.Login(ctx, tx, user)

	return responsedto.DefaultResponse{
		Message: "OK",
		Data:    user,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request requestdto.UserCreateRequest) responsedto.DefaultResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	// if err != nil {
	// 	panic(exception.NewNotFoundError(err.Error()))
	// }

	// user.Name = request.Name

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	user, _ = service.UserRepository.Register(ctx, tx, user)

	// return helper.ToCategoryResponse(user)
	return responsedto.DefaultResponse{
		Message: "OK",
		Data:    user,
	}
}
