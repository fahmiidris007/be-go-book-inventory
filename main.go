package main

import (
	"book-inventory/app"
	"book-inventory/auth"
	"book-inventory/db"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	connect := db.InitDB()

	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	handler := app.New(connect)
	//home
	router.GET("/", auth.HomeHandler)

	//login
	router.GET("/login", auth.LoginHandler)
	router.POST("/login", auth.LoginPostHandler)

	//get all books
	router.GET("/books", handler.GetBooks)

	router.Run(":5000")
}
