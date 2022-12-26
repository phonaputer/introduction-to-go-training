package cli

import (
	"customer_service/internal/customer/logic/cerr"
	"fmt"
)

func NewRootHandler(
	getHandler ErrorCommandHandler,
	createHandler ErrorCommandHandler,
) CommandHandler {
	result := &rootHandler{
		getHandler:    getHandler,
		createHandler: createHandler,
	}

	return withErrMsg(result)
}

type rootHandler struct {
	getHandler    ErrorCommandHandler
	createHandler ErrorCommandHandler
}

func (r *rootHandler) Handle(args []string) error {
	if len(args) < 1 {
		return cerr.NewUserMsg("not enough arguments")
	}

	switch args[0] {
	case "get":
		return r.getHandler.Handle(args[1:])
	case "create":
		return r.createHandler.Handle(args[1:])
	}

	return cerr.NewUserMsg(fmt.Sprintf("Unexpected subcommand: %s", args[0]))
}
