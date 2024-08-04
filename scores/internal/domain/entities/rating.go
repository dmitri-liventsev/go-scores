package entities

import (
	vo "go-scores/scores/internal/domain/value-objects"
	"time"
)

type Rating struct {
	ID         uint          `gorm:"primaryKey;autoIncrement"`
	Value      int           `gorm:"column:rating;not null"`
	TicketID   vo.TicketID   `gorm:"column:ticket_id;not null"`
	CategoryID vo.CategoryID `gorm:"column:rating_category_id;not null"`
	ReviewerID int           `gorm:"column:reviewer_id;not null"`
	RevieweeID int           `gorm:"column:reviewee_id;not null"`
	CreatedAt  time.Time     `gorm:"column:created_at;autoCreateTime"`
}

func (Rating) TableName() string {
	return "ratings"
}

func NewRating(id int, value int, categoryId int, ticketId int, createdAt time.Time) Rating {
	return Rating{
		ID:         uint(id),
		Value:      value,
		TicketID:   vo.NewTicketID(ticketId),
		CategoryID: vo.NewCategoryID(categoryId),
		CreatedAt:  createdAt,
	}
}
