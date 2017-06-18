package routers

import (
	"github.com/dung13890/go-services/gateway/handler"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	user := &handler.UserHandler{Service: handler.NewService("http://users:8080/api/v1/")}
	//product := &handler.ProductHandler{Service: handler.NewService("products:8080")}
	r := mux.NewRouter().StrictSlash(false)
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/users", user.Get).Methods("GET")
	// s.HandleFunc("/users", user.Create).Methods("POST")
	// s.HandleFunc("/users/{id}", user.Delete).Methods("DELETE")

	// s.HandleFunc("/products", user.Get).Methods("GET")
	// s.HandleFunc("/products", user.Create).Methods("POST")
	// s.HandleFunc("/products/{id}", user.Delete).Methods("DELETE")
	return s
}
