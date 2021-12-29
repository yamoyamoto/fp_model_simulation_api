package repository

type FPRepository interface{

}


type fpRepository struct{

}


func NewFPRepository() FPRepository {
	return &fpRepository{}
}