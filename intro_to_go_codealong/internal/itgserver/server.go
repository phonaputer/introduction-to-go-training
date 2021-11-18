package itgserver

import (
	"intro_to_go_codealong/internal/handler"
	"net/http"
)

func Run() {
	mux := http.NewServeMux()
	mux.Handle("/customers", http.HandlerFunc(handler.CustomerGetHandler))

	server := &http.Server{Addr: ":8080", Handler: mux}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

