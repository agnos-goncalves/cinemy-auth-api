package handlers

import (
	"cinemy-auth-api/repository"
	"cinemy-auth-api/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email     string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(context *gin.Context){
	var requestData LoginRequest

	if err:=context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err:=repository.UserRegister(requestData.Email, requestData.Password); if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered user"})
}

func LoginHandler(context *gin.Context) {

	var requestData LoginRequest

	if err:= context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.UserLogin(requestData.Email, requestData.Password)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} 

	payload := jwt.MapClaims{
		"sub": user.Id,
		"email": user.Email,
	}

	token, err:=utils.GenerateJWT(payload)

	if err!=nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, token)
}

func GuardAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty token"})
		c.Abort()
		return
	}

	if !strings.HasPrefix(tokenString, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid format token"})
		c.Abort() 
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secretKey := []byte("your-secret-key")
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		c.Abort()
		return
	}
	c.Next()
}







