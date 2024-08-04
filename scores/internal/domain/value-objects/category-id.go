package vo

import (
	"database/sql/driver"
	"errors"
)

type CategoryID int

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

func (c CategoryID) Value() (driver.Value, error) {
	return int(c), nil
}

func NewCategoryID(id int) CategoryID {
	return CategoryID(id)
}

func (c CategoryID) ToInt() int {
	return int(c)
}
