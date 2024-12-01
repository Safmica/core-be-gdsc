package controllers

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx, client *firestore.Client, ctx context.Context) error {
	return c.JSON(fiber.Map{
		"message": "Profile updated successfully",
	})
}