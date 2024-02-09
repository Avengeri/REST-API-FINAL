package model

type UserTodo struct {
	Id   int    `json:"id" swaggerignore:"true"`
	Age  int    `json:"age" example:"25"`
	Name string `json:"name" example:"Vitalya_realniy_keks"`
}
