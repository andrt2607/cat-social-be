package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	middleware.Handler.ServeHTTP(writer, request)
	// if "RAHASIA" == request.Header.Get("X-API-Key") {
	// 	// ok
	// } else {
	// 	// error
	// 	writer.Header().Set("Content-Type", "application/json")
	// 	writer.WriteHeader(http.StatusUnauthorized)

	// 	webResponse := responsedto.DefaultResponse{
	// 		// Code:   http.StatusUnauthorized,
	// 		// Status: "UNAUTHORIZED",
	// 	}

	// 	helper.WriteToResponseBody(writer, webResponse)
	// }
}
