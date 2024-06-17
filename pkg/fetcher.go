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
	station *models.Station,
	startDate string,
	wg *sync.WaitGroup,
	ch chan<- models.WeatherResultChannel,
	errCh chan<- error,
) {
	defer wg.Done()
	defer SendInfoLog("Finished fetching data for station " + strconv.Itoa(station.StationId))

	currentTime := time.Now()

	formattedDate := currentTime.Format("2006-01-02")

	apiUrl := GetApiUrl() + "?station=" + strconv.Itoa(station.StationId) + "&start=" + startDate + "&end=" + formattedDate

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

	ch <- models.WeatherResultChannel{
		StationName: station.StationName,
		Data:        weather,
	}
}

func FetchAllStationsData(startDate string) (map[string][]models.Weather, error) {
	stations := GetStations()

	var wg sync.WaitGroup

	weatherCh := make(chan models.WeatherResultChannel, len(stations))

	errCh := make(chan error, len(stations))

	for _, station := range stations {
		wg.Add(1)

		go FetchDailyWeatherData(station, startDate, &wg, weatherCh, errCh)
	}

	wg.Wait()

	close(weatherCh)
	close(errCh)

	weatherDataMap := make(map[string][]models.Weather)

	for result := range weatherCh {
		weatherDataMap[result.StationName] = append(weatherDataMap[result.StationName], result.Data...)
	}

	if len(errCh) > 0 {
		return nil, <-errCh
	}

	return weatherDataMap, nil
}
