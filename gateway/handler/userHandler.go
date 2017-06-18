package handler

import (
	"fmt"
	"github.com/dung13890/go-services/users/models"
	"io/ioutil"
	"net/http"
)

type UserHandler struct {
	Service *Service
}

func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {

	var target []models.Product

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://users:8080/api/v1/users", nil)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
