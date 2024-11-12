package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerNode interface {
	InitAndRun()
}

type serverNode struct {
	port string
}

func newServerNode(port string) ServerNode {
	return &serverNode{port: port}
}

func (s *serverNode) InitAndRun() {
	// Initialize a new Gin router
	router := gin.Default()

	// Define a simple GET route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start the server on port 8080
	router.Run(s.port)
}
