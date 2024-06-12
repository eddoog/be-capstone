package pkg

import (
	"io"
	"net/http"

	"github.com/eddoog/be-capstone/models"
)

func FetchDailyWeatherData(stationId string, startDate string, endDate string) ([]models.Weather, error) {
	apiUrl := GetApiUrl() + "?id=" + stationId + "&start=" + startDate + "&end=" + endDate

	resp, err := http.Get(apiUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	weather, err := models.MarshalWeather(body)

	if err != nil {
		return nil, err
	}

	return weather, nil
}
