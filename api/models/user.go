package models

import (
	"github.com/jinzhu/gorm"
)

// User interface
type User struct {
	gorm.Model
	UUID        string `json:"uuid" gorm:"type:varchar(100);unique_index"`
	AccessLevel int    `json:"access_level" valid:"int~access level is not correct.,required~access level required.,range(1|2)~access level must be equal to 1 or 2." gorm:"not_null"`
	FirstName   string `json:"first_name" valid:"required~first name is required.,length(2|100)~first name should be greater than 2."`
	LastName    string `json:"last_name" valid:"required~last name is required.,length(2|100)~last name should be greater than 2."`
	Email       string `json:"email" valid:"email~email is not correct.,required~email is required." gorm:"type:varchar(100);unique_index"`
	Password    string `json:"password" valid:"required~password is required." gorm:"not_null"`
	DateOfBirth string `json:"date_of_birth" valid:"rfc3339~date of birth is not correct.,required~date of birth is required."`
}
