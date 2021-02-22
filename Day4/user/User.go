package user

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Email string
	Password string
}

var UserDB = []User{}

var APISECRET = hex.EncodeToString([]byte("une-cle-de-32-bytes-de-long-là!"))

func signinSession(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	if user.Email != "" {
		c.String(http.StatusOK, "User successfully created")
	} else {
		c.String(http.StatusBadRequest, "Bad Request")
	}
}

func ApplyRoutes(r *gin.Engine) {
	r.POST("/signin-session", signinSession)
}