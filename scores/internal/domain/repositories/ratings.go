package repositories

import (
	"go-scores/scores/internal/domain/entities"
	"gorm.io/gorm"
	"time"
)

type Ratings struct {
	db *gorm.DB
}

// FindByPeriod find all entries between two dates.
func (repo *Ratings) FindByPeriod(start, end time.Time) ([]entities.Rating, error) {
	var result []entities.Rating
	startTime := start.Format("2006-01-02") + "T00:00:00"
	endTime := end.Format("2006-01-02") + "T23:59:59"

	err := repo.db.Where("created_at >= ? AND created_at <= ?", startTime, endTime).Find(&result).Error

	return result, err
}

// FindAll return all ratings.
func (repo *Ratings) FindAll() ([]entities.Rating, error) {
	var ratings []entities.Rating
	err := repo.db.Order("created_at asc").Find(&ratings).Error

	return ratings, err
}

func NewRatings(db *gorm.DB) Ratings {
	return Ratings{db}
}
