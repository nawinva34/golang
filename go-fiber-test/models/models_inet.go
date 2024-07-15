package models

// Body Parser method => Client send json to server
type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type UserRegister struct {
	Email        string `json:"email" validate:"required,email"`
	Username     string `json:"username" validate:"required,username_validate"`
	Password     string `json:"password" validate:"required,min=6,max=20"`
	LineID       string `json:"line_id" validate:"required"`
	Tel          string `json:"tel" validate:"required,tel"`
	BusinessType string `json:"business_type" validate:"required"`
	WebName      string `json:"web_name" validate:"required,min=2,max=30,web_validate"`
}
