package repository

import (
	"database/sql"
	"fmt"
)

type Customer struct {
	db *sql.DB
}

func NewCustomer(db *sql.DB) *Customer {
	return &Customer{
		db: db,
	}
}

func (c *Customer) GetOneByID(id int) (string, error) {
	row := c.db.QueryRow("SELECT id, first_name, middle_name, last_name, age FROM customers WHERE id=?", id)

	var resID, age int
	var firstName, lastName string
	var middleName *string

	err := row.Scan(&resID, &firstName, &middleName, &lastName, &age)
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
