package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dung13890/go-services/users/common"
	"github.com/dung13890/go-services/users/models"
	"github.com/dung13890/go-services/users/repo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type response struct {
	Message string        `json:"message,omitempty"`
	Status  int           `json:"status"`
	Item    *models.User  `json:"item,omitempty"`
	Items   []models.User `json:"items,omitempty"`
}

type UserHandler struct{}

func (u *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	c := context.DbCollection("users")
	defer context.Close()
	repo := &repo.UserRepo{c}
	user, err := repo.GetById(id)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	response := &response{
		Message: "success",
		Item:    &user,
		Status:  http.StatusOK,
	}
	rs, err := json.Marshal(response)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}

func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &repo.UserRepo{c}
	users := repo.GetAll()
	response := &response{
		Message: "success",
		Items:   users,
		Status:  http.StatusOK,
	}
	rs, err := json.Marshal(response)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &repo.UserRepo{c}
	err = repo.Create(&user)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	response := &response{
		Message: "success",
		Status:  http.StatusOK,
	}
	rs, err := json.Marshal(response)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}

func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &repo.UserRepo{c}
	err := repo.Delete(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			common.ResponseError(w, err, "Not found", http.StatusNotFound)
			return
		}
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	response := &response{
		Message: "success",
		Status:  http.StatusOK,
	}
	rs, err := json.Marshal(response)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}
