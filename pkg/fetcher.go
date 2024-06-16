package pkg

import (
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/eddoog/be-capstone/models"
)

func FetchDailyWeatherData(stationId string,
	startDate string,
	wg *sync.WaitGroup,
	ch chan<- []models.Weather,
	errCh chan<- error,
) {
	defer wg.Done()
	defer SendInfoLog("Finished fetching data for station " + stationId)

	currentTime := time.Now()

	formattedDate := currentTime.Format("2006-01-02")

	apiUrl := GetApiUrl() + "?id=" + stationId + "&start=" + startDate + "&end=" + formattedDate

	SendInfoLog("Fetching data from " + apiUrl)

	resp, err := http.Get(apiUrl)

	if err != nil {
		errCh <- err

		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		errCh <- err

		return
	}

	weather, err := models.MarshalWeather(body)

	if err != nil {
		errCh <- err

		return
	}

	ch <- weather
}
