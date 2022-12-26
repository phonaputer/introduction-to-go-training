package customer

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func RunCLI() {
	conf := loadConfig()

	logrus.SetLevel(conf.LogLevel)

	logrus.Trace("setting up CLI app")

	app, err := newCLIApp(conf)
	if err != nil {
		fmt.Println("An unexpected error has occurred!")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	logrus.Trace("setup complete")

	app.rootHandler.Handle(os.Args[1:])

	logrus.Trace("done. closing dependencies.")

	closeCLIApp(app)

	logrus.Trace("done. exiting.")
}
