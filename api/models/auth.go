package models

// Login struct
type Login struct {
	Login    string `json:"login" valid:"required~login is required."`
	Password string `json:"password" valid:"required~password is required."`
}
