package main

import (
	"os"

	"github.com/eddoog/be-capstone/api"
	"github.com/eddoog/be-capstone/pkg"
	"github.com/gofiber/fiber/v2"
)

func init() {
	pkg.LoadEnv()
	pkg.InitLog()
	pkg.NewConfig()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.SetupV1API(app)

	app.Listen(":" + os.Getenv("PORT"))
}
