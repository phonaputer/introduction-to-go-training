package handler

import (
	"intro_to_go_codealong/internal/domain"
	"intro_to_go_codealong/internal/itgerr"
	"intro_to_go_codealong/internal/view"
	"net/http"
	"strconv"
)

// interfaces needed by customer...
type CustomerValidator interface {
	CreateRequest(*http.Request) (*view.CustomerCreateReq, error)
}

type CustomerMapper interface {
	CreateReqToDomain(*view.CustomerCreateReq) *domain.Customer
}

type CustomerRepo interface {
	Create(customer *domain.Customer) error
	GetOneByID(id int) (string, error)
}

// customer handler struct
type Customer struct {
	customerValidator CustomerValidator
	customerMapper CustomerMapper
	customerRepo CustomerRepo
}

// initializer for handler.Customer. similar to a constructor.
func NewCustomer(customerValidator CustomerValidator, customerMapper CustomerMapper, customerRepo CustomerRepo) *Customer {
	return &Customer{
		customerValidator: customerValidator,
		customerMapper: customerMapper,
		customerRepo: customerRepo,
	}
}

// Create creates a new user
func (c *Customer) Create(w http.ResponseWriter, r *http.Request) error {
	reqView, err := c.customerValidator.CreateRequest(r)
	if err != nil {
		return err
	}

	customer := c.customerMapper.CreateReqToDomain(reqView)

	err = c.customerRepo.Create(customer)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)

	return nil
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
