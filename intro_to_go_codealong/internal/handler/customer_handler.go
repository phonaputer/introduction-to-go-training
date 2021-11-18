package handler

import (
	"intro_to_go_codealong/internal/itgerr"
	"net/http"
	"strconv"
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
func (c *Customer) GetOne(w http.ResponseWriter, r *http.Request) error {

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return itgerr.WithKind(err, itgerr.KindInvalidInput, "id is not valid")
	}

	userData, err := c.customerRepo.GetOneByID(id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userData))

	return nil
}
