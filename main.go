package main

import (
	"book-inventory/app"
	"book-inventory/auth"
	"book-inventory/db"
	"book-inventory/middleware"

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

	// get book by id
	router.GET("/book/:id", middleware.AuthValidation, handler.GetBookById)

	// add book
	router.GET("/addBook", middleware.AuthValidation, handler.AddBook)
	router.POST("/book", middleware.AuthValidation, handler.PostBook)

	// update book
	router.GET("/updateBook/:id", middleware.AuthValidation, handler.UpdateBook)
	router.POST("/updateBook/:id", middleware.AuthValidation, handler.PutBook)

	// delete book
	router.POST("/deleteBook/:id", middleware.AuthValidation, handler.DeleteBook)

	router.Run(":5000")
}
