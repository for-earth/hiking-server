package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kooku0/4earth/mountain"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/mountains", mountain.GetMountains)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
