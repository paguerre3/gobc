package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	common_api "github.com/paguerre3/blockchain/internal/common/api"
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
	log.Infof("Starting Blockchain Server on TCP Port %s", s.port)
	// Initialize a new Echo instance
	e := echo.New()

	// Define a simple GET route
	e.GET("/ping", common_api.Ping)

	// Start the server on the specified port
	e.Start(s.port)
}
