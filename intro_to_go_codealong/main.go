package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "docker:docker@tcp(127.0.0.1)/introtogo")
	if err != nil {
			panic(err)
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

	fmt.Println(result)
}

