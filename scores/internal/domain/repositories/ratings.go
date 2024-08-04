package repositories

import (
	"go-scores/scores/internal/domain/entities"
	"gorm.io/gorm"
	"time"
)

type Ratings struct {
	db *gorm.DB
}

func (repo *Ratings) GetByPeriod(start, end time.Time) ([]entities.Rating, error) {
	var result []entities.Rating
	startTime := start.Format("2006-01-02") + "T00:00:00"
	endTime := end.Format("2006-01-02") + "T23:59:59"

	err := repo.db.Where("created_at >= ? AND created_at <= ?", startTime, endTime).Find(&result).Error

	return result, err
}

func NewRatings(db *gorm.DB) Ratings {
	return Ratings{db}
}
