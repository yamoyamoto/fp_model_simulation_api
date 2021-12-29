package entity

// 結合定数の構造体
type Hebb struct{
	J [][]int64
}

// パターンの構造体
type Pattern struct{
	Body []int64
}


func (pattern *Pattern) CaluculateHebb(trainData *[][]int64) (Hebb){
	var hebb Hebb
	for i := range pattern.Body{
		var hebb_i []int64
		for j := range pattern.Body {
			var hebb_i_j int64 = 0
			for _, trainDataOne := range *trainData{
				hebb_i_j += trainDataOne[i] * trainDataOne[j]
			}
			hebb_i = append(hebb_i, hebb_i_j)
		}
		hebb.J = append(hebb.J, hebb_i)
	}
	return hebb
}