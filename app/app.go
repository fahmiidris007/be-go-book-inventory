package app

import (
	"book-inventory/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func New(db *gorm.DB) Handler {
	return Handler{DB: db}
}

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetBooks(c *gin.Context) {
	var books []models.Books

	h.DB.Find(&books)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Book Inventory",
		"payload": books,
		"auth":    c.Query("auth"),
	})
}

func (h *Handler) GetBookById(c *gin.Context) {
	var book models.Books
	id := c.Param("id")
	if h.DB.Find(&book, "id=?", id).RecordNotFound() {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.HTML(http.StatusOK, "book.html", gin.H{
		"title":   book.Title,
		"payload": book,
		"auth":    c.Query("auth"),
	})
}

func (h *Handler) AddBook(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{
		"title": "Add Book",
		"auth":  c.Query("auth"),
	})
}

func (h *Handler) PostBook(c *gin.Context) {
	var book models.Books
	err := c.Bind(&book)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	h.DB.Create(&book)
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/books?auth=%s", c.PostForm("auth")))
}

func (h *Handler) UpdateBook(c *gin.Context) {
	var books models.Books

	bookId := c.Param("id")
	if h.DB.Find(&books, "id=?", bookId).RecordNotFound() {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "not found"})
	}

	c.HTML(http.StatusOK, "form.html", gin.H{
		"title":   "add book",
		"payload": books,
		"auth":    c.Query("auth"),
	})
}

func (h *Handler) PutBook(c *gin.Context) {
	var books models.Books

	bookId := c.Param("id")
	if h.DB.Find(&books, "id=?", bookId).RecordNotFound() {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "not found"})
	}

	var reqBook = books
	c.Bind(&reqBook)

	h.DB.Model(&books).Where("id=?", bookId).Update(reqBook)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/book/%s?auth=%s", bookId, c.PostForm("auth")))
}

func (h *Handler) DeleteBook(c *gin.Context) {
	var books models.Books

	bookId := c.Param("id")
	h.DB.Delete(&books, "id=?", bookId)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/books?auth=%s", c.PostForm("auth")))
}
