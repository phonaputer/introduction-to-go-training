package logic

type CustomerService interface {
	Create(customer *Customer) (int, error)
	Get(id int) (*Customer, error)
}

type customerServiceImpl struct {
	customerRepository CustomerRepository
}

func (c *customerServiceImpl) Create(customer *Customer) (int, error) {
	return c.customerRepository.CreateCustomer(customer)
}

func (c *customerServiceImpl) Get(id int) (*Customer, error) {
	return c.customerRepository.GetCustomer(id)
}
