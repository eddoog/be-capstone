package api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Time(c *fiber.Ctx) error {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return c.SendStatus(500)
	}

	currentTime := time.Now().In(location)

	fmt.Println(currentTime)

	return c.SendString(currentTime.String())
}
