package pkg

import "sync"

var (
	globalConfig *Config
	configOnce   sync.Once
)

type Config struct {
	TimeWindow int
	ApiUrl     string
}

func GetConfig() *Config {
	configOnce.Do(InitializeConfig)

	return globalConfig
}

func InitializeConfig() {
	globalConfig = &Config{
		TimeWindow: 60,
		ApiUrl:     "https://d.meteostat.net/app/proxy/stations/daily",
	}
}

// GetTimeWindow returns the time window of the global configuration.
func GetTimeWindow() int {
	return GetConfig().TimeWindow
}

// GetApiUrl returns the API URL of the global configuration.
func GetApiUrl() string {
	return GetConfig().ApiUrl
}
