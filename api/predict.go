package api

import (
	"fmt"

	"github.com/eddoog/be-capstone/pkg"
	"github.com/gofiber/fiber/v2"
)

func Predict(c *fiber.Ctx) error {
	startDate, err := pkg.GetPastDate(pkg.GetTimeWindow())

	if err != nil {
		return c.Status(500).JSON(
			pkg.GetErrorMap(
				fmt.Errorf("error getting past date"), int16(500),
			),
		)
	}

	weathersData, err := pkg.FetchAllStationsData(startDate)

	if err != nil {
		return c.Status(500).JSON(pkg.GetErrorMap(
			fmt.Errorf("error fetching all stations data"), int16(500)))
	}

	stationMapTf, err := pkg.BuildStationMapTf(weathersData)

	if err != nil {
		return c.Status(500).JSON(pkg.GetErrorMap(
			fmt.Errorf("error building station map tensorflow"), int16(500)))
	}

	fmt.Println(stationMapTf)

	return c.Status(200).JSON(pkg.GetResponseMap(weathersData, int16(200)))
}
