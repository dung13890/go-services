package routers

import (
	"github.com/dung13890/go-services/users/handler"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	user := &handler.UserHandler{}
	r := mux.NewRouter().StrictSlash(false)
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/users", user.Get).Methods("GET")
	s.HandleFunc("/users", user.Create).Methods("POST")
	s.HandleFunc("/users/{id}", user.Delete).Methods("DELETE")
	return s
}
