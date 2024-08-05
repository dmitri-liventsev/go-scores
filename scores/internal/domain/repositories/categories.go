package repositories

import (
	"go-scores/scores/internal/domain/entities"
	"gorm.io/gorm"
)

// Categories repository.
type Categories struct {
	db *gorm.DB
}

// FindAll returns list of all categories.
func (repo *Categories) FindAll() ([]entities.Category, error) {
	var categories []entities.Category
	err := repo.db.Find(&categories).Error

	return categories, err
}

// NewCategories returns new Categories repository instance.
func NewCategories(db *gorm.DB) Categories {
	return Categories{db}
}
