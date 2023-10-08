package main

import (
	mail "cinemy-auth-api/communication"
	"cinemy-auth-api/handlers/authHandler"
	"cinemy-auth-api/middlewares/guards"
	"net/http"

	"github.com/gin-gonic/gin"
)

func buildRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	apiAuth := api.Group("/auth")
	apiPrivate := api.Group("/private")
	
	apiPrivate.Use(guards.Auth)
	
	api.GET("/", func(context *gin.Context){
		context.JSON(http.StatusOK, gin.H{ "message": "Welcome" })
	})
	apiPrivate.GET("/", func(context *gin.Context){
		mail.SendMailConfirmRegister("agnos.goncalves@gmail.com")
		context.JSON(http.StatusOK, gin.H{ "message": "Test Access" })
	})
	
	apiAuth.POST("/login", authHandler.Login)
	apiAuth.POST("/register", authHandler.Register)
	apiAuth.PUT("/register/confirm", authHandler.ConfirmRegister)
	return router
}

func main(){
	router := buildRouter()
	router.Run(":8080")
}