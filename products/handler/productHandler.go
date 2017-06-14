package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dung13890/go-services/products/common"
	"github.com/dung13890/go-services/products/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type response struct {
	Message string           `json:"message,omitempty"`
	Status  int              `json:"status"`
	Item    *models.Product  `json:"item,omitempty"`
	Items   []models.Product `json:"items,omitempty"`
}

type ProductHandler struct{}

func (p *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("products")
	repo := &models.ProductRepo{c}
	products := repo.GetAll()
	response := &response{
		Message: "success",
		Items:   products,
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

func (p *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		common.ResponseError(w, err, "Invalid Product data", 500)
		return
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("products")
	repo := &models.ProductRepo{c}
	err = repo.Create(&product)
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

func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("products")
	repo := &models.ProductRepo{c}
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
