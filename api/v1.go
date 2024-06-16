package api

import "github.com/gofiber/fiber/v2"

func SetupV1API(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Post("/predict", Predict)
	v1.Get("/time", Time)
}
