package entity

// 結合定数の構造体
type Hebb struct {
	J [][]int64
}

// パターンの構造体
type Pattern struct {
	Body []int64
}

func (pattern *Pattern) CaluculateHebb(trainData *[][]int64) Hebb {
	var hebb Hebb
	for i := range pattern.Body {
		var hebb_i []int64
		for j := range pattern.Body {
			var hebb_i_j int64 = 0
			for _, trainDataOne := range *trainData {
				hebb_i_j += trainDataOne[i] * trainDataOne[j]
			}
			hebb_i = append(hebb_i, hebb_i_j)
		}
		hebb.J = append(hebb.J, hebb_i)
	}
	return hebb
}

func (pattern *Pattern) ExecDynamics(hebb *Hebb, dynamicsCount int) {
	for i := 0; i < dynamicsCount; i++ {
		var next_S_ij int64 = 0
		for j := 0; j < len(pattern.Body); j++ {
			for k := 0; k < len(pattern.Body); k++ {
				next_S_ij += hebb.J[i][j] * pattern.Body[i]
			}
		}
		if next_S_ij >= 0 {
			pattern.Body[i] = 1
		} else {
			pattern.Body[i] = -1
		}
	}
}
