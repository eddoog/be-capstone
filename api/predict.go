package api

import (
	"github.com/eddoog/be-capstone/pkg"
	"github.com/gofiber/fiber/v2"
)

func Predict(c *fiber.Ctx) error {
	stations := pkg.GetStations()

	return c.SendString("Predict")
}
