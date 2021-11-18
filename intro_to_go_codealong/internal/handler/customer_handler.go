package handler

import (
	"fmt"
	"intro_to_go_codealong/internal/repository"
	"net/http"
)

func CustomerGetHandler(w http.ResponseWriter, r *http.Request) {
	userData, err := repository.GetCustomer()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userData))
}
