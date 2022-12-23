package logic

import "errors"

var ErrNotFound = errors.New("not found")

type CustomerRepository interface {
	CreateCustomer(c *Customer) (int, error)
	GetCustomer(id int) (*Customer, error)
}
