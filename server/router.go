package server

import "github.com/gin-gonic/gin"

func registerHandler(s *Server) {
	r := s.engine

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
