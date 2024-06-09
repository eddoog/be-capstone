package main

import (
	"os"

	"github.com/eddoog/be-capstone/pkg"
	"github.com/gofiber/fiber/v2"
)

func init() {
	pkg.InitLog()
	pkg.LoadEnv()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + os.Getenv("PORT"))
}
