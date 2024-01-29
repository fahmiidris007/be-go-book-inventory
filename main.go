package main

import (
	"book-inventory/app"
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
	router.GET("/", handler.GetBooks)

	router.Run(":5000")
}
