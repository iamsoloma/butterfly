package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iamsoloma/butterfly"
	"github.com/iamsoloma/butterfly/system"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9060"
		fmt.Println("Can`t parse port! Used standard: 1106.")
	}

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		timeout = int(time.Duration.Seconds(60))
		fmt.Println("Can`t parse timeout! Used standard: 60 sec.")
	}

	bodyLimit, err := strconv.Atoi(os.Getenv("BODYLIMIT"))
	if err != nil {
		bodyLimit = 1024 * 1024 * 1024 * 1024
		fmt.Println("Can`t parse BodyLimit! Used standard: 1024*1024*1024*1024.")
	}

	s := NewServer(":" + port, bodyLimit, time.Duration(timeout))
	log.Fatal(s.Start())
}

func (s *Server) Start() error {
	f := fiber.New(
		fiber.Config{
			BodyLimit:         s.bodyLimit,
			IdleTimeout:       s.idleTimeout,
			Prefork:           false,
			StreamRequestBody: true,
		},
	)

	//main
	f.Get("/health", s.Health)

	return f.Listen(s.listenAddr)
}

func (s *Server) Health(c *fiber.Ctx) (err error) {
	memory := system.ReadMemoryStats()

	resp := butterfly.Health{
		Status:           "ok",
		UTC:              time.Now().UTC().String(),
		NodeType:         "stone",
		Version:          "0.1.0",
		TotalMemory:     memory.MemTotal,
		AvailableMemory: memory.MemAvailable,
	}
	return c.JSON(resp)
}
