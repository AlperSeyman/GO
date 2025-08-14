package models

import "gorm.io/gorm"

/*

gorm.Model provides a predefined struct named gorm.Model, which inclued commonly used fields:
gorm.Model definition

type Model struct {
ID        uint           `gorm:"primaryKey"`
CreatedAt time.Time
UpdatedAt time.Time
DeletedAt gorm.DeletedAt `gorm:"index"`
}

*/

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"` // pointer to a string; allowing for null values
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}
