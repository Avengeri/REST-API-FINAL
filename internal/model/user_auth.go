package model

type UserAuth struct {
	Id          int    `json:"id" swaggerignore:"true"`
	Username    string `json:"username" binding:"required" example:"vitalya-XAKEP"`
	Password    string `json:"password" binding:"required" example:"qwerty"`
	Email       string `json:"email" binding:"required" example:"vitalya@mail.ru"`
	CreatedTime string `json:"created_time" swaggerignore:"true"`
}
