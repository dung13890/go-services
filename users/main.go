package main

import (
	"log"
	"net/http"

	"github.com/dung13890/go-services/users/common"
	"github.com/dung13890/go-services/users/routers"
)

func main() {
	common.Init()
	router := routers.InitRoutes()
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
