package main

import (
	"github.com/yamoto0628/fp_model_sumilation_api/controller"
	"github.com/yamoto0628/fp_model_sumilation_api/model/repository"
	"net/http"
)

var tr = repository.NewFPRepository()
var tc = controller.NewFPController(tr)
var ro = controller.NewRouter(tc)

func main() {
	http.HandleFunc("/", ro.HandleFPRequest)
	_ = http.ListenAndServe(":8001", nil)
}
