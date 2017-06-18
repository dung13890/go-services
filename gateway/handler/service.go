package handler

import (
	"log"
	"net/http"
)

type Service struct {
	Url string
}

func NewService(url string) *Service {
	service := &Service{
		Url: url,
	}
	return service
}

func (s *Service) Request(method string, endpoint string, callback func(*http.Request)) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(method, s.Url+endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	callback(req)
	resp, _ := client.Do(req)
	return resp
}
