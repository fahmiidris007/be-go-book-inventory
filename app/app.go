package app

import (
	"book-inventory/models"
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
	})
}
