package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yamoto0628/fp_model_simulation_api/controller"
	"github.com/yamoto0628/fp_model_simulation_api/model/repository"
)

var tr = repository.NewFPRepository()
var tc = controller.NewFPController(tr)
var ro = controller.NewRouter(tc)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("server run on port:%v", port)
	http.HandleFunc("/api", ro.HandleFPRequest)
	_ = http.ListenAndServe(":"+port, nil)
}
