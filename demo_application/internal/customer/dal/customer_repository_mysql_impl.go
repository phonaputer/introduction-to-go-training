package dal

import (
	logic2 "customer_service/internal/customer/logic"
	"database/sql"
	"errors"
	"fmt"
)

func NewCustomerRepositoryMySQLImpl(
	db *sql.DB,
) logic2.CustomerRepository {
	return &customerRepositoryMySQLImpl{
		db: db,
	}
}

type customerRepositoryMySQLImpl struct {
	db *sql.DB
}

func (c *customerRepositoryMySQLImpl) CreateCustomer(customer *logic2.Customer) (int, error) {
	const query = `INSERT INTO customers 
    						(first_name, middle_name, last_name, age)
    					VALUES 
    					    (?, ?, ?, ?)`

	result, err := c.db.Exec(query, customer.FirstName, customer.MiddleName, customer.LastName, customer.Age)
	if err != nil {
		return 0, fmt.Errorf("INSERT customer: %w", err)
	}

	newCustomerID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get new customer ID: %w", err)
	}

	return int(newCustomerID), nil
}

func (c *customerRepositoryMySQLImpl) GetCustomer(id int) (*logic2.Customer, error) {
	const query = `SELECT first_name, middle_name, last_name, age FROM customers WHERE id=?`

	row := c.db.QueryRow(query, id)

	result := logic2.Customer{ID: id}

	err := row.Scan(&result.FirstName, &result.MiddleName, &result.LastName, &result.Age)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, logic2.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("SELECT customer: %w", err)
	}

	return &result, nil
}
