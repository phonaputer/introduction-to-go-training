package customer

import (
	"customer_app/internal/customer/cli"
	"customer_app/internal/customer/dal"
	"customer_app/internal/customer/logic"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type cliApp struct {
	db *sql.DB

	customerRepository logic.CustomerRepository
	customerService    logic.CustomerService

	commandValidator cli.CustomerCommandValidator
	getHandler       cli.ErrorCommandHandler
	createHandler    cli.ErrorCommandHandler
	rootHandler      cli.CommandHandler
}

func newCLIApp(conf *appConfig) (*cliApp, error) {
	var app cliApp
	var err error

	app.db, err = dal.NewMySQLClient(&dal.MySQLClientConfig{
		User:      conf.MySQL.User,
		Password:  conf.MySQL.Password,
		Address:   conf.MySQL.Address,
		DB:        conf.MySQL.DB,
		TimeoutMs: conf.MySQL.TimeoutMs,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting MySQL connection: %w", err)
	}

	app.customerRepository = dal.NewCustomerRepositoryMySQLImpl(app.db)
	app.customerService = logic.NewCustomerService(app.customerRepository)

	app.commandValidator = cli.NewCustomerCommandValidator()
	app.getHandler = cli.NewGetHandler(app.customerService, app.commandValidator)
	app.createHandler = cli.NewCreateHandler(app.customerService, app.commandValidator)

	app.rootHandler = cli.NewRootHandler(app.getHandler, app.createHandler)

	return &app, nil
}

func closeCLIApp(app *cliApp) {
	err := app.db.Close()
	if err != nil {
		logrus.WithError(err).Debug("error closing DB pool")
	}
}
