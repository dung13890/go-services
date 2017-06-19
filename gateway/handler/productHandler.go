package handler

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"

	"github.com/dung13890/go-services/gateway/common"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Service *Service
}

type Product struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Sku         string        `json:"sku,omitempty"`
	Price       int           `json:"price,omitempty"`
	Description string        `json:"description,omitempty"`
	Date        time.Time     `json:"date, omitempty"`
}

type ProductResponse struct {
	Message string    `json:"message,omitempty"`
	Status  int       `json:"status"`
	Item    *Product  `json:"item,omitempty"`
	Items   []Product `json:"items,omitempty"`
}

func (p *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	target := ProductResponse{}
	resp := p.Service.Request("GET", "products/"+id, func(*http.Request) {})
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}

	rs, err := json.Marshal(target)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}

func (p *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	target := ProductResponse{}
	resp := p.Service.Request("GET", "products", func(*http.Request) {})
	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	rs, err := json.Marshal(target)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}

func (p *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	target := ProductResponse{}
	client := &http.Client{}
	req, err := http.NewRequest("POST", p.Service.Url+"products", r.Body)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	rs, err := json.Marshal(target)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}

func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	target := ProductResponse{}
	resp := p.Service.Request("DELETE", "products/"+id, func(*http.Request) {})
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}

	rs, err := json.Marshal(target)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}
