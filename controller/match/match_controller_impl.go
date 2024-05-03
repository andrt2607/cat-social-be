package controller

import (
	"cat-social-be/helper"
	requestdto "cat-social-be/model/dto/request"
<<<<<<< Updated upstream
	repository "cat-social-be/repository/match"
=======
	matchRepository "cat-social-be/repository/match"
>>>>>>> Stashed changes
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleInternalServerError(c *gin.Context, err error) {
	fmt.Println("Internal Server Error:", err)
	// Mengirim respons dengan status HTTP 500 (Internal Server Error)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal Server Error",
	})
}

<<<<<<< Updated upstream
=======
// func GetMatchs(c *gin.Context) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			handleInternalServerError(c, fmt.Errorf("%v", err))
// 		}
// 	}()
// 	db := c.MustGet("db").(*sql.DB)
// 	//call repository
// 	catRepository.GetCats(c, db)
// }

>>>>>>> Stashed changes
func CreateMatch(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
<<<<<<< Updated upstream
	userLoginRequest := requestdto.UserLoginRequest{}
	c.ShouldBindJSON(&userLoginRequest)
	if err := helper.ValidateStruct(&userLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if !repository.IsEmailExist(c, db, userLoginRequest.Email) {
	// 	c.JSON(http.StatusConflict, gin.H{
	// 		"error": "User Not Found",
	// 	})
	// 	return
	// }

	repository.GetMatches(c, db)
}

//	func GetMatches(c *gin.Context) {
//		defer func() {
//			if err := recover(); err != nil {
//				handleInternalServerError(c, fmt.Errorf("%v", err))
//			}
//		}()
//		db := c.MustGet("db").(*sql.DB)
//		userCreateRequest := requestdto.UserCreateRequest{}
//		c.ShouldBindJSON(&userCreateRequest)
//		if err := helper.ValidateStruct(&userCreateRequest); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		if repository.IsEmailExist(c, db, userCreateRequest.Email) {
//			c.JSON(http.StatusConflict, gin.H{
//				"error": "Email already exist",
//			})
//			return
//		}
//		hashedPassword, err := helper.HashPassword(userCreateRequest.Password)
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		userCreateRequest.Password = hashedPassword
//		registerResponse, _ := repository.Register(c, db, userCreateRequest)
//		c.JSON(http.StatusCreated, registerResponse)
//	}
func GetMatches(c *gin.Context) {
=======
	matchCreateRequest := requestdto.MatchCreateRequest{}
	c.ShouldBindJSON(&matchCreateRequest)
	//validasi input
	if err := helper.ValidateStruct(&matchCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validasi Match Request
	createRequestReady, err_code, err_message := matchRepository.ValidateCreateMatch(c, db, matchCreateRequest)
	if err_code != nil {
		c.JSON(err_code, gin.H{
			"error": err_message,
		})
		return
	}

	//call repository
	matchCreateResponse, _ := matchRepository.CreateMatch(c, db, createRequestReady)
	c.JSON(http.StatusCreated, matchCreateResponse)
}

func ApproveMatch(c *gin.Context) {
>>>>>>> Stashed changes
	defer func() {
		if err := recover(); err != nil {
			handleInternalServerError(c, fmt.Errorf("%v", err))
		}
	}()
	db := c.MustGet("db").(*sql.DB)
<<<<<<< Updated upstream
	repository.GetMatches(c, db)
}

func ApproveMatch(c *gin.Context) {

}

func RejectMatch(c *gin.Context) {

}

func DeleteMatch(c *gin.Context) {

}
=======
	matchApproveRequest := requestdto.MatchApproveRequest{}
	c.ShouldBindJSON(&matchApproveRequest)
	//validasi input
	if err := helper.ValidateStruct(&matchApproveRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//validasi Match Request
	_, err_code, err_message := matchRepository.ValidateApproveMatch(c, db, matchApproveRequest)
	if err_code != nil {
		c.JSON(err_code, gin.H{
			"error": err_message,
		})
		return
	}

	//call repository
	matchApproveResponse, _ := matchRepository.ApproveMatch(c, db, matchApproveRequest)
	c.JSON(http.StatusCreated, matchApproveResponse)
}

// func DeleteCat(c *gin.Context) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			handleInternalServerError(c, fmt.Errorf("%v", err))
// 		}
// 	}()
// 	db := c.MustGet("db").(*sql.DB)
// 	catCreateRequest := requestdto.CatCreateRequest{}
// 	c.ShouldBindJSON(&catCreateRequest)
// 	//validasi id cat
// 	_, err := catRepository.FindCatById(c, db)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "Id cat is not found",
// 		})
// 		return

// 	}
// 	//call repository
// 	catCreateResponse, _ := catRepository.DeleteCat(c, db, catCreateRequest)
// 	c.JSON(http.StatusCreated, catCreateResponse)
// }
>>>>>>> Stashed changes
