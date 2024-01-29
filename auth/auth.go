package auth

import (
	"book-inventory/models"
	"net/http"
	"net/url"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
	})
}

func LoginPostHandler(c *gin.Context) {
	var credential models.Login
	err := c.Bind(&credential)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "Invalid request",
		})
	}
	if credential.Username != models.USER || credential.Password != models.PASSWORD {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"content": "Invalid username or password",
		})
	} else {
		claim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
			Issuer:    "book inventory",
			IssuedAt:  time.Now().Unix(),
		}
		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		token, err := sign.SignedString([]byte(models.SECRET_KEY))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"content": "Invalid request",
			})
			c.Abort()
		}
		q := url.Values{}
		q.Set("auth", token)
		location := url.URL{
			Path:     "/books",
			RawQuery: q.Encode(),
		}
		c.Redirect(http.StatusMovedPermanently, location.RequestURI())
	}
}
