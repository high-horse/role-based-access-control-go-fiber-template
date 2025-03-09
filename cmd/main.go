package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"rbac/db/config"
	"rbac/db/pool"
	"rbac/internal/router"

	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	confg := config.LoadConfig()
	err := pool.ConnectDB(confg)
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	defer pool.DisconnectDB()

	serve()
}

// unused function (dont use this funcction )
func waitsignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down the server...")
}

func serve() {
	log.Println("starting the server ...")
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: false,
		// ColorScheme: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "role-permission-server v1.0.1",
	})
	app.Use(logger.New())
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	router.AuthRoutes(app)
	// List all registered routes
	for _, routes := range app.Stack() {
		// Iterate over individual routes in the group
		for _, route := range routes {
			fmt.Printf("Method: %s\tPath: %s\n", route.Method, route.Path)
		}
	}
	log.Fatal(app.Listen(":8000"))
}
