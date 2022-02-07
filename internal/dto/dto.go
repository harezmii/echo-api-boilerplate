package dto

type UserDto struct {
	Id     int    `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Email  string `json:"email" xml:"email"`
	Status *bool  `json:"status" xml:"status"`
}
