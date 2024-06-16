package pkg

import (
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/eddoog/be-capstone/models"
)

func FetchDailyWeatherData(
	stationId string,
	startDate string,
	wg *sync.WaitGroup,
	ch chan<- []models.Weather,
	errCh chan<- error,
) {
	defer wg.Done()
	defer SendInfoLog("Finished fetching data for station " + stationId)

	currentTime := time.Now()

	formattedDate := currentTime.Format("2006-01-02")

	apiUrl := GetApiUrl() + "?station=" + stationId + "&start=" + startDate + "&end=" + formattedDate

	SendInfoLog("Fetching data from " + apiUrl)

	resp, err := http.Get(apiUrl)

	if err != nil {
		SendWarnLog("Error fetching data: " + err.Error())
		errCh <- err

		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		SendWarnLog("Error reading response body: " + err.Error())
		errCh <- err

		return
	}

	weather, err := models.MarshalWeather(body)

	if err != nil {
		SendWarnLog("Error marshalling weather data: " + err.Error())
		errCh <- err

		return
	}

	ch <- weather
}

func FetchAllStationsData(startDate string) ([]models.Weather, error) {
	stations := GetStations()

	var wg sync.WaitGroup

	weatherCh := make(chan []models.Weather, len(stations))

	errCh := make(chan error, len(stations))

	for _, station := range stations {
		wg.Add(1)

		go FetchDailyWeatherData(strconv.Itoa(station.StationId), startDate, &wg, weatherCh, errCh)
	}

	wg.Wait()

	close(weatherCh)
	close(errCh)

	var allWeather []models.Weather

	for weather := range weatherCh {
		allWeather = append(allWeather, weather...)
	}

	if len(errCh) > 0 {
		return nil, <-errCh
	}

	return allWeather, nil
}
