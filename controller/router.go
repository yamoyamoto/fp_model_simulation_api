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
	switch r.Method {
	case "POST":
		err := ro.tc.FPSumilation(w, r)
		if err != nil {
			fmt.Printf("FPSumilation error: %s", err)
		}
	default:
		w.WriteHeader(405)
	}
}
