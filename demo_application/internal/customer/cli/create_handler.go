package cli

import (
	"customer_service/internal/customer/logic"
	"fmt"
)

func NewCreateHandler(
	customerService logic.CustomerService,
	commandValidator CustomerCommandValidator,
) ErrorCommandHandler {
	return &createHandler{
		customerService:  customerService,
		commandValidator: commandValidator,
	}
}

type createHandler struct {
	customerService  logic.CustomerService
	commandValidator CustomerCommandValidator
}

func (c *createHandler) Handle(args []string) error {
	// 1. validate input
	inputData, err := c.commandValidator.ValidateCreate(args)
	if err != nil {
		return fmt.Errorf("validate create data: %w", err)
	}

	newCustomer := logic.Customer{
		FirstName:  *inputData.FirstName,
		MiddleName: inputData.MiddleName,
		LastName:   *inputData.LastName,
		Age:        *inputData.Age,
	}

	// 2. create customer
	id, err := c.customerService.Create(&newCustomer)
	if err != nil {
		return fmt.Errorf("create customer: %w", err)
	}

	// 3. print new customer's ID
	fmt.Printf("New Customer created with ID: %v\n", id)

	return nil
}
