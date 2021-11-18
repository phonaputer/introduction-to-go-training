package handler

import (
	"fmt"
	"net/http"
)

// interfaces needed by customer...
type CustomerRepo interface {
	GetOneByID(id int) (string, error)
}

// customer handler struct
type Customer struct {
	customerRepo CustomerRepo
}

// initializer for handler.Customer. similar to a constructor.
func NewCustomer(customerRepo CustomerRepo) *Customer {
	return &Customer{
		customerRepo: customerRepo,
	}
}

// GetOne gets the data from one customer
func (c *Customer) GetOne(w http.ResponseWriter, r *http.Request) {
	userData, err := c.customerRepo.GetOneByID(11)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userData))
}
