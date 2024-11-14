package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
	// Initialize a new Echo instance
	e := echo.New()

	// Define a simple GET route
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	// Start the server on the specified port
	e.Start(s.port)
}
