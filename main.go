package main

import(
	"fmt"
	"net/http"
)

func handler(web http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(web, "hello world!!")
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}