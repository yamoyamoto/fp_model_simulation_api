package controller

import (
	"fmt"
	"net/http"
)

type Router interface {
	HandleFPRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc FPController
}

// Routerのコンストラクタ。引数にFPControllerを受け取り、Router構造体のポインタを返却する。
func NewRouter(tc FPController) Router {
	return &router{tc}
}

func (ro *router) HandleFPRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		err := ro.tc.FPSimulation(w, r)
		if err != nil {
			fmt.Printf("FPSimulation error: %s", err)
		}
	case "OPTIONS":
		// for CORS
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	default:
		w.WriteHeader(405)
	}
}
