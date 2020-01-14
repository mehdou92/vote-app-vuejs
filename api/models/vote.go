package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Vote structure
type Vote struct {
	gorm.Model
	UUID        string         `json:"uuid" gorm:"type:varchar(100);unique_index"`
	Title       string         `json:"title" valid:"required~title is required.,length(2|100)~description should be greater than 2"`
	Description string         `json:"description" valid:"required~description is required.,length(5|255)~description should be greater than 5 and lower than 255."`
	UUIDVote    pq.StringArray `json:"uuid_votes" gorm:"type:text[]"`
	StartDate   string         `json:"start_date" valid:"rfc3339~start date is not correct."`
	EndDate     string         `json:"end_date" valid:"rfc3339~end date is not correct."`
}
