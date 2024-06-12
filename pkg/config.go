package pkg

type Config struct {
	TimeWindow int
	ApiUrl     string
}

func NewConfig() *Config {
	return &Config{
		TimeWindow: 60,
		ApiUrl:     "https://d.meteostat.net/app/proxy/stations/hourly",
	}
}
