package main

import (
	"github.com/yamoto0628/fp_model_sumilation_api/controller"
	"github.com/yamoto0628/fp_model_sumilation_api/model/repository"
	"net/http"
	"os"
)

var tr = repository.NewFPRepository()
var tc = controller.NewFPController(tr)
var ro = controller.NewRouter(tc)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.HandleFunc("/api", ro.HandleFPRequest)
	_ = http.ListenAndServe(port, nil)
}
