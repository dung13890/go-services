package routers

import (
	"github.com/dung13890/go-services/products/handler"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	product := &handler.ProductHandler{}
	r := mux.NewRouter().StrictSlash(false)
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/products/{id}", product.GetById).Methods("GET")
	s.HandleFunc("/products", product.Get).Methods("GET")
	s.HandleFunc("/products", product.Create).Methods("POST")
	s.HandleFunc("/products/{id}", product.Delete).Methods("DELETE")
	return s
}
