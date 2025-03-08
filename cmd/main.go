package main

import (
	"log"
	"os"
	"os/signal"
	"rpba-app/db/config"
	"rpba-app/db/pool"
	"rpba-app/internal/router"

	"syscall"

	"github.com/gofiber/fiber/v2"
)



func main() {
	confg := config.LoadConfig()
	err := pool.ConnectDB(confg)
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	defer pool.DisconnectDb()

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
	    Prefork:       false,
	    CaseSensitive: false,
	    // ColorScheme: true,
	    StrictRouting: true,
	    ServerHeader:  "Fiber",
	    AppName: "role-permission-server v1.0.1",
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	router.AuthRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
