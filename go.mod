module github.com/yamoto0628/fp_model_sumilation_api

go 1.17

replace github.com/yamoto0628/fp_model_sumilation_api/model/repository => ./model/repository

replace github.com/yamoto0628/fp_model_sumilation_api/model/entity => ./model/entity

replace github.com/yamoto0628/fp_model_sumilation_api/controller/dto => ./controller/dto

require (
	github.com/yamoto0628/fp_model_sumilation_api/controller/dto v0.0.0-00010101000000-000000000000
	github.com/yamoto0628/fp_model_sumilation_api/model/entity v0.0.0-00010101000000-000000000000
	github.com/yamoto0628/fp_model_sumilation_api/model/repository v0.0.0-00010101000000-000000000000
)
