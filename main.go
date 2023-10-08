package main

import (
	"cinemy-auth-api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func buildRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	apiAuth := api.Group("/auth")
	apiPrivate := api.Group("/private")
	
	apiPrivate.Use(handlers.GuardAuth)
	
	api.GET("/", func(context *gin.Context){
		context.JSON(http.StatusOK, gin.H{ "message": "Welcome" })
	})
	apiPrivate.GET("/", func(context *gin.Context){
		context.JSON(http.StatusOK, gin.H{ "message": "Test Access" })
	})
	
	apiAuth.POST("/login", handlers.LoginHandler)
	apiAuth.POST("/register", handlers.RegisterHandler)
	return router
}

func main(){
	router := buildRouter()
	router.Run(":8080")
}