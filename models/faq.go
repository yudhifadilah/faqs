package models

import "gorm.io/gorm"

type FAQ struct {
	gorm.Model
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
