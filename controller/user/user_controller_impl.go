package controller

import (
	"cat-social-be/helper"
	requestdto "cat-social-be/model/dto/request"
	responsedto "cat-social-be/model/dto/response"
	service "cat-social-be/service/user"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(UserService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: UserService,
	}
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Print("masuk login")
	userCreateRequest := requestdto.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	categoryResponse := controller.UserService.Login(request.Context(), userCreateRequest)
	resultResponse := responsedto.DefaultResponse{
		Message: "OK",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(writer, resultResponse)
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := requestdto.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	registerResponse := controller.UserService.Register(request.Context(), userCreateRequest)
	resultResponse := responsedto.DefaultResponse{
		Message: "Boleh",
		Data:    registerResponse,
	}

	helper.WriteToResponseBody(writer, resultResponse)
}
