package dto

type FPRequest struct {
	TrainData   [][]int64 `json:"train_data"`
	InputPattern []int64  `json:"input_pattern"`
	DynamicsCount int64   `json:"dynamics_count"`
}

type FPResponse struct {
	OutputPattern []int64 `json:"output_pattern"`
}
