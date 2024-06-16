package api

import (
	"fmt"

	"github.com/eddoog/be-capstone/pkg"
	"github.com/gofiber/fiber/v2"
)

func Predict(c *fiber.Ctx) error {
	startDate, err := pkg.GetPastDate(pkg.GetTimeWindow())

	if err != nil {
		return c.Status(500).SendString("Error getting past date")
	}

	weathersData, err := pkg.FetchAllStationsData(startDate)

	if err != nil {
		return c.Status(500).SendString("Error fetching weathers data")
	}

	fmt.Println(weathersData)

	return c.SendString("Predict")
}
