package server

import "github.com/gin-gonic/gin"

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	s := &Server{
		engine: r,
	}
	registerHandler(s)
	return s
}

func (s *Server) Run() {
	s.engine.Run()
}
