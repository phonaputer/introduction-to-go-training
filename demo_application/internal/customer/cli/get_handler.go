package cli

import (
	"customer_service/internal/customer/logic"
	"fmt"
)

func NewGetHandler(
	customerService logic.CustomerService,
	commandValidator CustomerCommandValidator,
) ErrorCommandHandler {
	return &getHandler{
		customerService:  customerService,
		commandValidator: commandValidator,
	}
}

type getHandler struct {
	customerService  logic.CustomerService
	commandValidator CustomerCommandValidator
}

func (g *getHandler) Handle(args []string) error {
	// 1. validate input
	id, err := g.commandValidator.ValidateGet(args)
	if err != nil {
		return fmt.Errorf("validating args: %w", err)
	}

	// 2. fetch customer
	customer, err := g.customerService.Get(id)
	if err != nil {
		return fmt.Errorf("fetch customer: %w", err)
	}

	// 3. print customer
	fmt.Println(fmt.Sprintf("Customer %v: ", customer.ID))
	fmt.Println(fmt.Sprintf("First Name: %s", customer.FirstName))

	if customer.MiddleName != nil {
		fmt.Println(fmt.Sprintf("Middle Name: %s", *customer.MiddleName))
	}

	fmt.Println(fmt.Sprintf("Last Name: %s", customer.LastName))
	fmt.Println(fmt.Sprintf("Age: %v", customer.Age))

	return nil
}
