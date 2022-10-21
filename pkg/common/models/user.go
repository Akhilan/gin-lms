package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // adds ID, created_at etc.
	Name       string `json:"name"`
	Place      string `json:"place"`
	Mail       string `json:"mail"`
}
