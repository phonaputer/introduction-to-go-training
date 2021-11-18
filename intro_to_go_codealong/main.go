package main

import (
	"net/http"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/customers", http.HandlerFunc(pocHandler))

	server := &http.Server{Addr: ":8080", Handler: mux}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func pocHandler(w http.ResponseWriter, r *http.Request) {
	userData := getFromMySQL()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userData))
}

func getFromMySQL() string {
	db, err := sql.Open("mysql", "docker:docker@tcp(127.0.0.1)/introtogo")
	if err != nil {
			panic(err)
	}
	defer db.Close()

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

	return result
}

