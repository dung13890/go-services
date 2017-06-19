package handler

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"

	"github.com/dung13890/go-services/gateway/common"
	"github.com/gorilla/mux"
)

type User struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string        `json:"name,omitempty"`
	Email string        `json:"email,omitempty"`
	Date  time.Time     `json:"date, omitempty"`
}

type UserResponse struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status"`
	Item    *User  `json:"item,omitempty"`
	Items   []User `json:"items,omitempty"`
}

type UserHandler struct {
	Service *Service
}

func (u *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	target := UserResponse{}
	resp := u.Service.Request("GET", "users/"+id, func(*http.Request) {})
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

func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	target := UserResponse{}
	resp := u.Service.Request("GET", "users", func(*http.Request) {})
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

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	target := UserResponse{}
	client := &http.Client{}
	req, err := http.NewRequest("POST", u.Service.Url+"users", r.Body)
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

func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	target := UserResponse{}
	resp := u.Service.Request("DELETE", "users/"+id, func(*http.Request) {})
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
