package main

import (
	"api/service"
	"api/system"
	"log"
	"net/http"
)

const (
	port = "3001"
)

func main() {
	service := service.NewService()
	r := system.SetupRouter(service)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
