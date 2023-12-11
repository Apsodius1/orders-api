package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list order")
}