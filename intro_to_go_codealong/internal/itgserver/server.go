package itgserver

import (
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"intro_to_go_codealong/internal/client"
	"intro_to_go_codealong/internal/handler"
	"intro_to_go_codealong/internal/mapper"
	"intro_to_go_codealong/internal/repository"
	"intro_to_go_codealong/internal/validator"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	deps, err := injectDependencies()
	if err != nil {
		logrus.WithError(err).Fatal("error creating deps")
	}

	server := &http.Server{Addr: ":8080", Handler: setupPathMatching(deps)}

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	go runServer(server)

	recievedSig := <-signalChan

	logrus.WithField("signal", recievedSig).Info("stopping for signal")

	err = server.Shutdown(context.Background())
	if err != nil {
		logrus.WithError(err).Error("error shutting down server")
	}

	err = closeDependencies(deps)
	if err != nil {
		logrus.WithError(err).Error("error closing deps")
	}
}

func runServer(s *http.Server) {
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("error running server")
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

	customerHandler := handler.NewCustomer(&validator.Customer{}, &mapper.Customer{}, customerRepo)

	return &dependencies{
		db:              db,
		customerRepo:    customerRepo,
		customerHandler: customerHandler,
	}, nil
}

func closeDependencies(deps *dependencies) error {
	return deps.db.Close()
}

func setupPathMatching(deps *dependencies) http.Handler {
	r := mux.NewRouter()
	r.Handle("/customers", handler.ErrRespAdaptor(deps.customerHandler.GetOne)).Methods("GET")
	r.Handle("/customers", handler.ErrRespAdaptor(deps.customerHandler.Create)).Methods("POST")

	return r
}
