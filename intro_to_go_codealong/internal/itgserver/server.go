package itgserver

import (
	"database/sql"
	"intro_to_go_codealong/internal/client"
	"intro_to_go_codealong/internal/handler"
	"intro_to_go_codealong/internal/repository"
	"net/http"
)

func Run() {
	deps, err := injectDependencies()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/customers", http.HandlerFunc(deps.customerHandler.GetOne))

	server := &http.Server{Addr: ":8080", Handler: mux}

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

type dependencies struct {
	db *sql.DB

	customerRepo *repository.Customer

	customerHandler *handler.Customer
}

func injectDependencies() (*dependencies, error) {
	db, err := client.GetMySQLDB()
	if err != nil {
		return nil, err
	}

	customerRepo := repository.NewCustomer(db)

	customerHandler := handler.NewCustomer(customerRepo)

	return &dependencies{
		db:             db,
		customerRepo:    customerRepo,
		customerHandler: customerHandler,
	}, nil
}

