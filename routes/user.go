package routes

import (
	"context"
	"core-be-gdsc/controllers"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes (app *fiber.App, client *firestore.Client, ctx context.Context){
	app.Post("/user", func(c *fiber.Ctx) error {
		return controllers.CreateUser(c, client, ctx)
	})
}