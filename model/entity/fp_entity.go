package entity

import (
	"math"
	"math/rand"
)

type Hebb [][]int64

type Pattern []int64

type PatternFromDynamics struct {
	Count   int     `json:"count"`
	Pattern Pattern `json:"pattern"`
}

func CalculateHebb(pattern []int64, trainData *[][]int64) Hebb {
	var hebb Hebb
	N := len(*trainData)
	for i := range pattern {
		var hebb_i []int64
		for j := range pattern {
			var hebb_i_j int64 = 0
			for _, trainDataOne := range *trainData {
				hebb_i_j += trainDataOne[i] * trainDataOne[j]
			}
			hebb_i = append(hebb_i, hebb_i_j/int64(N))
		}
		hebb = append(hebb, hebb_i)
	}
	return hebb
}

func (hebb Hebb) ExecDynamics(pattern []int64, dynamicsCount int, beta float64) []PatternFromDynamics {
	var patternFromDynamicsAll []PatternFromDynamics
	patternLen := len(pattern)
	phaseInterval := dynamicsCount / 10
	for count := 0; count < dynamicsCount; count++ {
		var next_pattern []int64

		for i := 0; i < len(pattern); i++ {
			var next_S_ij int64 = 0

			for j := 0; j < len(pattern); j++ {
				if i == j {
					continue
				}
				next_S_ij = next_S_ij + hebb[i][j]*pattern[j]
			}

			next_pattern = append(next_pattern, decideNextS_ij(next_S_ij, beta))
			// fmt.Printf("i=%vにおいて、next_S_ijは%vなので、次のS_iは、%vです\n", i, next_S_ij, next_pattern[i])
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
