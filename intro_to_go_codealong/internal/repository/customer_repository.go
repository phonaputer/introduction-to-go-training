package repository

import (
	"fmt"
	"intro_to_go_codealong/internal/client"
)

func GetCustomer() (string, error) {
	db, err := client.GetMySQLDB()
	if err != nil {
		return "", err
	}

	row := db.QueryRow("SELECT id, first_name, middle_name, last_name, age FROM customers WHERE id=?", 11)

	var id, age int
	var firstName, lastName string
	var middleName *string

	err = row.Scan(&id, &firstName, &middleName, &lastName, &age)
	if err != nil {
		panic(err)
	}

	result := fmt.Sprintf("%v: %s ", id, firstName)

	if middleName != nil {
		result += *middleName + " "
	}

	result += fmt.Sprintf("%s. Age: %v ", lastName, age)

	return result, nil
}
