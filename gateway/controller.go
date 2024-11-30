package main

import "net/http"

type controller struct {
}

func NewController() *controller {
	return &controller{}
}

func (c *controller) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/customers/{customerId}/orders", c.CreateOrder)
}

func (c *controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
}
