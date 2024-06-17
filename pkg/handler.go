package pkg

import "github.com/gofiber/fiber/v2"

func GetErrorMap(err error, status int16) fiber.Map {
	return fiber.Map{
		"error":  err.Error(),
		"status": status,
		"data":   nil,
	}
}

func GetResponseMap(data interface{}, status int16) fiber.Map {
	return fiber.Map{
		"error":  nil,
		"status": status,
		"data":   data,
	}
}
