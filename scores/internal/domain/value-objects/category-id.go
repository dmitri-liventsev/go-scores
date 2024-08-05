package vo

import (
	"database/sql/driver"
	"errors"
)

// CategoryID category primary key
type CategoryID int

// Scan implementation of Scanner interface.
func (c *CategoryID) Scan(value interface{}) error {
	if value == nil {
		*c = 0
		return nil
	}
	val, ok := value.(int64)
	if !ok {
		return errors.New("failed to scan CategoryID")
	}
	*c = CategoryID(val)
	return nil
}

// Value implementation of Valuer interface.
func (c CategoryID) Value() (driver.Value, error) {
	return int(c), nil
}

// ToInt converting value object to int value.
func (c CategoryID) ToInt() int {
	return int(c)
}

// NewCategoryID returns new CategoryID instance.
func NewCategoryID(id int) CategoryID {
	return CategoryID(id)
}
