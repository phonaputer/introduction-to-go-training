package repository

import (
	"database/sql"
	"fmt"
	"intro_to_go_codealong/internal/domain"
	"intro_to_go_codealong/internal/itgerr"
)

type Customer struct {
	db *sql.DB
}

func NewCustomer(db *sql.DB) *Customer {
	return &Customer{
		db: db,
	}
}

func (c *Customer) Create(customer *domain.Customer) error {
	_, err := c.db.Exec("INSERT INTO customers (first_name, middle_name, last_name, age) VALUES (?, ?, ?, ?)",
		customer.FirstName,
		customer.MiddleName,
		customer.LastName,
		customer.Age)

	if err != nil {
		return fmt.Errorf("error inserting customer: %w", err)
	}

	return nil
}

func (c *Customer) GetOneByID(id int) (string, error) {
	row := c.db.QueryRow("SELECT id, first_name, middle_name, last_name, age FROM customers WHERE id=?", id)

	var resID, age int
	var firstName, lastName string
	var middleName *string

	err := row.Scan(&resID, &firstName, &middleName, &lastName, &age)
	if err == sql.ErrNoRows {
		return "", itgerr.WithKind(err, itgerr.KindNotFound, "customer not found")
	}
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v: %s ", resID, firstName)

	if middleName != nil {
		result += *middleName + " "
	}

	result += fmt.Sprintf("%s. Age: %v ", lastName, age)

	return result, nil
}
