package main

import (
	"net/http"

	"github.com/ojoaobronstrup/i_prime/controller"
)

func main() {
	http.HandleFunc("POST /login", controller.ValidateUser)

	http.ListenAndServe(":8080", nil)
}
