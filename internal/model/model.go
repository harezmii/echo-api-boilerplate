package model

type User struct {
	Name     string `json:"name" xml:"name" form:"name" validate:"required,min=3,max=30"`
	Email    string `json:"email" xml:"email" form:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
	Status   *bool  `json:"status" xml:"status" form:"status" validate:"required"`
}
