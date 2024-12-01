package main

import (
	"context"
	"core-be-gdsc/database"
	"core-be-gdsc/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.EnvInit()
	client := database.ConnectFirestore()
	app := fiber.New()
	ctx := context.Background()

	routes.UserRoutes(app, client, ctx)

	app.Listen(":3000")
}