package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yamoto0628/fp_model_sumilation_api/controller/dto"
	"github.com/yamoto0628/fp_model_sumilation_api/model/entity"
	"github.com/yamoto0628/fp_model_sumilation_api/model/repository"
)

type FPController interface {
	FPSumilation(w http.ResponseWriter, r *http.Request) error
}
type fPController struct {
	tr repository.FPRepository
}

func NewFPController(tr repository.FPRepository) FPController {
	return &fPController{tr}
}

// シュミレーションフロー
func (tc *fPController) FPSumilation(w http.ResponseWriter, r *http.Request) error{
	// INPUTを受け取る
	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil {
		return err
	}
	var fpRequest dto.FPRequest
	err = json.Unmarshal(body, &fpRequest)
	if err != nil{
		return err
	}

	pattern := entity.Pattern{Body: fpRequest.InputPattern}
	trainData := fpRequest.TrainData

	hebb := pattern.CaluculateHebb(&trainData)
	pattern.ExecDynamics(&hebb, 100)

	// JSONに変換
	output, _ := json.MarshalIndent(pattern, "", "\t\t")
	fmt.Print(output)

	// JSONを返却
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		return err
	}
	return nil
}
