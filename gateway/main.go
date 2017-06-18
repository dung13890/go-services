package main

import (
	"log"
	"net/http"

	"github.com/dung13890/go-services/gateway/routers"
)

func main() {
	router := routers.InitRoutes()
	server := &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
