package main

import (
	"SoftwareGoDay2/routes"
	"SoftwareGoDay2/server"
	"fmt"
)

func main() {
	s := server.NewServer()
    routes.ApplyRoutes(s.Router)
	err := s.Router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
}
