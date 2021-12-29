package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yamoto0628/fp_model_sumilation_api/controller/dto"
	"github.com/yamoto0628/fp_model_sumilation_api/model/entity"
	"github.com/yamoto0628/fp_model_sumilation_api/model/repository"
)

type FPController interface{
	FPSumilation(w http.ResponseWriter, r *http.Request)
}
type fPController struct{
	tr repository.FPRepository
}
func NewFPController(tr repository.FPRepository) FPController {
	return &fPController{tr}
}

// シュミレーションフロー
func (tc *fPController) FPSumilation(w http.ResponseWriter, r *http.Request) {
	// INPUTを受け取る
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var fpRequest dto.FPRequest
	json.Unmarshal(body, &fpRequest)

	pattern := entity.Pattern{Body: fpRequest.InputPattern}
	trainData := fpRequest.TrainData

	hebb := pattern.CaluculateHebb(&trainData)

	// JSONに変換
	output, _ := json.MarshalIndent(hebb.J, "", "\t\t")
	fmt.Print(output)

	// ダミーフロー
	var dammy []int64
	dammy = append(dammy, 1)
	dammy = append(dammy, 2)
	dammy = append(dammy, 3)
	dammyResponce, _ := json.Marshal(dammy)

	// JSONを返却
	w.Header().Set("Content-Type", "application/json")
	w.Write(dammyResponce)
}