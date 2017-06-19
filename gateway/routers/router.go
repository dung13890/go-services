package routers

import (
	"github.com/dung13890/go-services/gateway/handler"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	user := &handler.UserHandler{Service: handler.NewService("http://users:8080/api/v1/")}
	product := &handler.ProductHandler{Service: handler.NewService("http://products:8080/api/v1/")}
	order := &handler.OrderHandler{Service: handler.NewService("http://orders:8080/api/v1/")}
	r := mux.NewRouter().StrictSlash(false)
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/users/{id}", user.GetById).Methods("GET")
	s.HandleFunc("/users", user.Get).Methods("GET")
	s.HandleFunc("/users", user.Create).Methods("POST")
	s.HandleFunc("/users/{id}", user.Delete).Methods("DELETE")

	s.HandleFunc("/products/{id}", product.GetById).Methods("GET")
	s.HandleFunc("/products", product.Get).Methods("GET")
	s.HandleFunc("/products", product.Create).Methods("POST")
	s.HandleFunc("/products/{id}", product.Delete).Methods("DELETE")

	s.HandleFunc("/orders/{user}", order.GetUser).Methods("GET")
	s.HandleFunc("/orders", order.Create).Methods("POST")
	s.HandleFunc("/orders/{id}", order.Delete).Methods("DELETE")
	return s
}
