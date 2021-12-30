package main

import(
	// "fmt"
	"net/http"
	"github.com/yamoto0628/fp_model_sumilation_api/controller"
	"github.com/yamoto0628/fp_model_sumilation_api/model/repository"
)

var tr = repository.NewFPRepository()
var tc = controller.NewFPController(tr)
var ro = controller.NewRouter(tc)

func main(){
	http.HandleFunc("/", ro.HandleFPRequest)
	_ = http.ListenAndServe(":8001", nil)
}