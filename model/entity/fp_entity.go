package entity

import (
	"math"
	"math/rand"
)

type J [][]int64

type Pattern []int64

type PatternFromDynamics struct {
	Count   int     `json:"count"`
	Pattern Pattern `json:"pattern"`
}

func MakeJ(pattern []int64, trainData *[][]int64) J {
	var J J
	N := len(*trainData)
	for i := range pattern {
		var J_i []int64
		for j := range pattern {
			var J_ij int64 = 0
			for _, trainDataOne := range *trainData {
				J_ij += trainDataOne[i] * trainDataOne[j]
			}
			J_i = append(J_i, J_ij/int64(N))
		}
		J = append(J, J_i)
	}
	return J
}

func (J *J) ExecDynamics(pattern []int64, dynamicsCount int, beta float64) []PatternFromDynamics {
	var patternFromDynamicsAll []PatternFromDynamics
	patternLen := len(pattern)
	phaseInterval := dynamicsCount / 10

	nowPattern := make([]int64, patternLen)
	copy(nowPattern, pattern)
	for count := 0; count < dynamicsCount; count++ {
		var next_pattern []int64

		for i := 0; i < len(nowPattern); i++ {
			var h_ij int64 = 0

			for j := 0; j < len(nowPattern); j++ {
				if i == j {
					continue
				}
				h_ij = h_ij + (*J)[i][j]*nowPattern[j]
			}

			next_pattern = append(next_pattern, decideNextS_ij(h_ij, beta))
		}

		if phaseInterval == 0 || count%phaseInterval == 0 {
			patternCopy := make([]int64, patternLen)
			copy(patternCopy, next_pattern)
			patternFromDynamics := PatternFromDynamics{
				Count:   count,
				Pattern: patternCopy,
			}
			patternFromDynamicsAll = append(patternFromDynamicsAll, patternFromDynamics)
		}
		print("\n")

		nowPattern = next_pattern
	}
	return patternFromDynamicsAll
}

func decideNextS_ij(h_ij int64, beta float64) int64 {
	r := rand.Float64()
	if r <= W(h_ij, beta) {
		return 1
	} else {
		return -1
	}
}

func W(h_ij int64, beta float64) float64 {
	return 1 / (1 + math.Exp(-2*beta*float64(h_ij)))
}
