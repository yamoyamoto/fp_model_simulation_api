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
func (tc *fPController) FPSumilation(w http.ResponseWriter, r *http.Request) error {
	var fpRequest dto.FPRequest
	err := json.NewDecoder(r.Body).Decode(&fpRequest)
	if err != nil {
		fmt.Print("json parse error\n")
		return err
	}

	trainData := fpRequest.TrainData

	hebb := entity.CaluculateHebb(fpRequest.InputPattern, &trainData)
	var outputStruct dto.FPResponse
	outputStruct.OutputPattern = hebb.ExecDynamics(fpRequest.InputPattern, int(fpRequest.DynamicsCount))
	outputStruct.Hebb = hebb

	output, _ := json.Marshal(outputStruct)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(output)
	if err != nil {
		return err
	}
	return nil
}
