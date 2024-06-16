package pkg

import (
	"time"
)

func GetTodayDate() (string, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		SendInfoLog("Error loading location: " + err.Error())

		return "", err
	}

	currentTime := time.Now().In(location)

	return currentTime.Format("2006-01-02"), nil
}

func GetPastDate(days int) (string, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		SendWarnLog("Error loading location: " + err.Error())

		return "", err
	}

	currentTime := time.Now().In(location)

	pastTime := currentTime.AddDate(0, 0, -days)

	return pastTime.Format("2006-01-02"), nil
}
