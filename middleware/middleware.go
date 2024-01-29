package middleware

import (
	"book-inventory/models"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthValidation(c *gin.Context) {
	var tokenString string
	tokenString = c.Query("auth")
	if tokenString == "" {
		tokenString = c.PostForm("auth")
		if tokenString == "" {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"content": "token login not found",
			})
			c.Abort()
		}
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, invalid := token.Method.(*jwt.SigningMethodHMAC); !invalid {
			return nil, fmt.Errorf("Invalid token ", token.Header["alg"])
		}
		return []byte(models.SECRET_KEY), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
		c.Next()
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "token is expiry"})
		c.Abort()
	}

}
