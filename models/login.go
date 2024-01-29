package models

type Login struct {
	ID       int    `json:"id" form:"id" gorm:"primaryKey"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

const (
	USER       = "admin"
	PASSWORD   = "1234"
	SECRET_KEY = "secret"
)
