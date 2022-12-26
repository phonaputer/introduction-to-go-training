package logic

import (
	"customer_service/internal/customer/logic/cerr"
	"errors"
	"fmt"
)

type CustomerService interface {
	Create(customer *Customer) (int, error)
	Get(id int) (*Customer, error)
}

func NewCustomerService(
	customerRepository CustomerRepository,
) CustomerService {
	return &customerServiceImpl{
		customerRepository: customerRepository,
	}
}

type customerServiceImpl struct {
	customerRepository CustomerRepository
}

func (c *customerServiceImpl) Create(customer *Customer) (int, error) {
	return c.customerRepository.CreateCustomer(customer)
}

func (c *customerServiceImpl) Get(id int) (*Customer, error) {
	customer, err := c.customerRepository.GetCustomer(id)
	if errors.Is(err, ErrNotFound) {
		return nil, cerr.SetUserMsg(err, "customer not found")
	}
	if err != nil {
		return nil, fmt.Errorf("get customer: %w", err)
	}

	return customer, nil
}
