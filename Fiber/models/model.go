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
	Name    string
	Company *string // pointer to a string; allowing for null values
	Email   *string
	Phone   int
}
