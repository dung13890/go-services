package routers

import (
	"github.com/dung13890/go-services/orders/handler"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	order := &handler.OrderHandler{}
	r := mux.NewRouter().StrictSlash(false)
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/orders/{user}", order.GetUser).Methods("GET")
	s.HandleFunc("/orders", order.Create).Methods("POST")
	s.HandleFunc("/orders/{id}", order.Delete).Methods("DELETE")
	return s
}
