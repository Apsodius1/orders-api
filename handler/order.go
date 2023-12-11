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

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get order by id")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update order by id")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete order by id")
}