package routes

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func health(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func world(c *gin.Context) {
	c.JSON(200, gin.H{"message": "world"})
}

func repeatMyQuery(c *gin.Context) {
	query := c.Query("a")
	c.JSON(200, gin.H{"message": query})
}

func repeatMyParam(c *gin.Context) {
	message := c.Param("message")
	if len(message) > 0 {
		c.String(http.StatusOK, message)
	} else {
		c.String(http.StatusBadRequest, message)
	}
}

func repeatMyBody(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	c.JSON(200, gin.H{"message": body})
}

func repeatMyHeader(c *gin.Context) {
	head := c.Request.Header
	c.JSON(200, gin.H{"message": head})
}

func repeatMyCookie(c *gin.Context) {
	cookies := c.Request.Cookies
	c.JSON(200, gin.H{"message": cookies})
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/health", health)
	r.GET("/hello", world)
	r.GET("/repeat-my-query", repeatMyQuery)
	r.GET("/repeat-my-param/:message", repeatMyParam)
	r.POST("/repeat-my-body", repeatMyBody)
	r.GET("/repeat-my-header", repeatMyHeader)
	r.GET("/repeat-my-cookie", repeatMyCookie)
}