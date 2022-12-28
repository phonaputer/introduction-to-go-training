package cli

import (
	"customer_app/internal/customer/logic"
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
	fmt.Printf("Customer %v:\n", customer.ID)
	fmt.Printf("First Name: %s\n", customer.FirstName)

	if customer.MiddleName != nil {
		fmt.Printf("Middle Name: %s\n", *customer.MiddleName)
	}

	fmt.Printf("Last Name: %s\n", customer.LastName)
	fmt.Printf("Age: %v\n", customer.Age)

	return nil
}
