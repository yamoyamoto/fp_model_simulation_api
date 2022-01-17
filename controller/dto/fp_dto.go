package dto

import (
	"github.com/yamoto0628/fp_model_simulation_api/model/entity"
)

type FPRequest struct {
	TrainData     [][]int64 `json:"train_data"`
	InputPattern  []int64   `json:"input_pattern"`
	DynamicsCount int64     `json:"dynamics_count"`
	Beta          float64   `json:"beta"`
}

type FPResponse struct {
	OutputPattern []entity.PatternFromDynamics `json:"body"`
	Hebb          entity.J                     `json:"hebb"`
}
