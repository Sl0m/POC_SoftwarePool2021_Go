package server

import "github.com/gin-gonic/gin"

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	var s *Server = new(Server)
	s.Router = gin.Default()
	return s
}
