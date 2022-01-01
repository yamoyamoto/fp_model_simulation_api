package entity

type Hebb [][]int64

type Pattern []int64

type PatternFromDynamics struct {
	Count   int     `json:"count"`
	Pattern Pattern `json:"pattern"`
}

func CaluculateHebb(pattern []int64, trainData *[][]int64) Hebb {
	var hebb Hebb
	for i := range pattern {
		var hebb_i []int64
		for j := range pattern {
			var hebb_i_j int64 = 0
			for _, trainDataOne := range *trainData {
				hebb_i_j += trainDataOne[i] * trainDataOne[j]
			}
			hebb_i = append(hebb_i, hebb_i_j)
		}
		hebb = append(hebb, hebb_i)
	}
	return hebb
}

func (hebb Hebb) ExecDynamics(pattern []int64, dynamicsCount int) []PatternFromDynamics {
	var patternFromDynamicsAll []PatternFromDynamics
	fhaseInterval := dynamicsCount / 5
	for i := 0; i < dynamicsCount; i++ {
		if i%fhaseInterval == 0 {
			PatternFromDynamics := PatternFromDynamics{
				Count:   i,
				Pattern: pattern,
			}
			patternFromDynamicsAll = append(patternFromDynamicsAll, PatternFromDynamics)
		}
		var next_S_ij int64 = 0
		for j := 0; j < len(pattern); j++ {
			for k := 0; k < len(pattern); k++ {
				next_S_ij += hebb[j][k] * pattern[j]
			}
			if next_S_ij >= 0 {
				pattern[j] = 1
			} else {
				pattern[j] = -1
			}
		}
	}
	return patternFromDynamicsAll
}
