package controller

import(
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"../model/repository"
)


// 外部に公開するインターフェイス
type FPController interface{
	FPSumilation(w http.ResponseWriter, r *http.Request)
}


// 非公開のTodoController構造体
type FpController struct {
	tr repository.FPRepository
}


// TodoControllerのコンストラクタ。
// 引数にTodoRepositoryを受け取り、TodoController構造体のポインタを返却する。
func NewFPController(tr repository.FPRepository) FPController {
	return &FpController{tr}
}


// シュミレーションフロー
func (tc *FpController) FPSumilation(w http.ResponseWriter, r *http.Request) {


	// JSONに変換
	// output, _ := json.MarshalIndent(todosResponse.Todos, "", "\t\t")

	// JSONを返却
	w.Header().Set("Content-Type", "application/json")
	// w.Write(output)
}

