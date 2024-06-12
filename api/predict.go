package api

import "github.com/gofiber/fiber/v2"

func Predict(c *fiber.Ctx) error {
	return c.SendString("Predict")
}
