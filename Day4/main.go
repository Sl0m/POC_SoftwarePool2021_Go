package main

import (
	"SoftwareGoDay4/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
    user.ApplyRoutes(s)
	err := s.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
}
