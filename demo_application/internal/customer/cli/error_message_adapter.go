package cli

import (
	"customer_service/internal/customer/logic/cerr"
	"fmt"
	"github.com/sirupsen/logrus"
)

type ErrorCommandHandler interface {
	Handle(args []string) error
}

func withErrMsg(next ErrorCommandHandler) CommandHandler {
	return &errorCommandHandlerImpl{
		next: next,
	}
}

type errorCommandHandlerImpl struct {
	next ErrorCommandHandler
}

func (e *errorCommandHandlerImpl) Handle(args []string) {
	err := e.next.Handle(args)
	if err != nil {
		printErrorMessage(err)
	}
}

func printErrorMessage(err error) {
	fmt.Println("An error occurred while executing the Customer CLI")

	errMsg := "An unexpected error occurred!"
	if userMsg, ok := cerr.GetUserMsg(err); ok {
		errMsg = userMsg
	}
	fmt.Println("Error: " + errMsg)

	fmt.Println("Usage: ")
	fmt.Println("./customers get ${CUSTOMER_ID}")
	fmt.Println("./customers create -f=${PATH_TO_CUSTOMER_JSON_FILE}")

	logrus.WithError(err).Debug("Error handling command")
}
