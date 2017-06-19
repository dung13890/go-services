package handler

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"

	"github.com/dung13890/go-services/gateway/common"
	"github.com/gorilla/mux"
)

type OrderHandler struct {
	Service *Service
}

type OrderItem struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	User     *User         `json:"user,omitempty"`
	Products []string      `json:"products,omitempty"`
}

type OrderResponse struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status"`
	Item    *struct {
		Id       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
		User     string        `json:"user,omitempty"`
		Products []string      `json:"products,omitempty"`
		Date     time.Time     `json:"date, omitempty"`
	} `json:"item, omitempty"`
}

type OrderInformationResponse struct {
	Message string     `json:"message,omitempty"`
	Status  int        `json:"status"`
	Item    *OrderItem `json:"item,omitempty"`
}

func (o *OrderHandler) getUserById(id string) *User {
	target := UserResponse{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://users:8080/api/v1/users/"+id, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		log.Fatal(err)
	}
	return target.Item
}

func (o *OrderHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	target := OrderResponse{}
	resp := o.Service.Request("GET", "orders/"+user, func(*http.Request) {})
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	response := &OrderInformationResponse{
		Message: "success",
		Status:  http.StatusOK,
	}
	fmt.Println(target.Item)
	if target.Item != nil {
		obj := o.getUserById(user)
		response.Item = &OrderItem{
			Id:       target.Item.Id,
			User:     obj,
			Products: target.Item.Products,
		}
	}
	if err != nil {
		common.ResponseError(w, err, "Invalid User data", 500)
		return
	}
	rs, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rs)
}

func (o *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	target := OrderResponse{}
	client := &http.Client{}
	req, err := http.NewRequest("POST", o.Service.Url+"orders", r.Body)
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

func (o *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	target := OrderResponse{}
	resp := o.Service.Request("DELETE", "orders/"+id, func(*http.Request) {})
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
