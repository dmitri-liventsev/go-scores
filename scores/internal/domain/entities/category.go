package entities

import vo "go-scores/scores/internal/domain/value-objects"

// Category ratings category.
type Category struct {
	ID     vo.CategoryID `gorm:"primaryKey;autoIncrement"`
	Name   string        `gorm:"not null"`
	Weight float32       `gorm:"not null"`
}

// TableName implementation of Tabler interface.
func (Category) TableName() string {
	return "rating_categories"
}
