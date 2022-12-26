package customer

import "github.com/sirupsen/logrus"

type appConfig struct {
	LogLevel logrus.Level
	MySQL    mysql
}

type mysql struct {
	User      string
	Password  string
	Address   string
	DB        string
	TimeoutMs int
}

func loadConfig() *appConfig {
	// FIXME read this from a file instead of hardcoding

	return &appConfig{
		LogLevel: logrus.DebugLevel,
		MySQL: mysql{
			User:      "docker",
			Password:  "docker",
			Address:   "localhost:3306",
			DB:        "introtogo",
			TimeoutMs: 5000,
		},
	}
}
