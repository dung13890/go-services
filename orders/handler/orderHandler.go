package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dung13890/go-services/orders/common"
	"github.com/dung13890/go-services/orders/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type response struct {
	Message string        `json:"message,omitempty"`
	Status  int           `json:"status"`
	Item    *models.Order `json:"item,omitempty"`
}

type OrderHandler struct{}

func (o *OrderHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["user"]
	context := NewContext()
	c := context.DbCollection("orders")
	defer context.Close()
	repo := &models.OrderRepo{c}
	order, err := repo.GetUser(id)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	response := &response{
		Message: "success",
		Item:    &order,
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

func (o *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	order := models.Order{}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		common.ResponseError(w, err, "Invalid Product data", 500)
		return
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("orders")
	repo := &models.OrderRepo{c}
	err = repo.Create(&order)
	if err != nil {
		common.ResponseError(w, err, "Invalid Product data", 500)
		return
	}
	response := &response{
		Message: "success",
		Status:  http.StatusOK,
	}
	rs, err := json.Marshal(response)
	if err != nil {
		common.ResponseError(w, err, "Invalid Product data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}

func (o *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("orders")
	repo := &models.OrderRepo{c}
	err := repo.Delete(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			common.ResponseError(w, err, "Not found", http.StatusNotFound)
			return
		}
		common.ResponseError(w, err, "Invalid Product data", 500)
		return
	}
	response := &response{
		Message: "success",
		Status:  http.StatusOK,
	}
	rs, err := json.Marshal(response)
	if err != nil {
		common.ResponseError(w, err, "Invalid Product data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}
