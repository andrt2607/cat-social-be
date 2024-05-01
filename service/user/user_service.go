package service

import (
	requestdto "cat-social-be/model/dto/request"
	responsedto "cat-social-be/model/dto/response"
	"context"
)

type UserService interface {
	Login(ctx context.Context, request requestdto.UserCreateRequest) responsedto.DefaultResponse
	Register(ctx context.Context, request requestdto.UserCreateRequest) responsedto.DefaultResponse
}
