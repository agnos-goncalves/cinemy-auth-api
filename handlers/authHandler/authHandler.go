package authHandler

import (
	"cinemy-auth-api/service/userService"
	"cinemy-auth-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequestDTO struct {
	Email     string `json:"email"`
	Password string `json:"password"`
}

type LoginRequestDTO struct {
	Email     string `json:"email"`
	Password string `json:"password"`
}

type ConfirmRegisterDTO struct {
	Token string `json:"token"`
}

type PasswordChangeDTO struct {
	Email       string `json:"email"`
	CurrentPass string `json:"current_password"`
	NewPass     string `json:"new_password"`
}

func Register(context *gin.Context){
	var requestData RegisterRequestDTO

	if err:=context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err:= userService.Register(requestData.Email, requestData.Password); if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered user"})
}

func Login(context *gin.Context) {

	var requestData LoginRequestDTO

	if err:= context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err:=utils.IsValidEmail(requestData.Email); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := userService.Login(requestData.Email, requestData.Password)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} 

	context.JSON(http.StatusOK, token)
}

func ConfirmRegister(context *gin.Context){
	var requestData ConfirmRegisterDTO

	if err:= context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userService.Active(requestData.Token)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}

func PasswordChange(context *gin.Context) { 

	var requestData PasswordChangeDTO

	if err:= context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_,err := userService.PasswordChange(requestData.Email, requestData.CurrentPass, requestData.NewPass)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "password changed"})
}






