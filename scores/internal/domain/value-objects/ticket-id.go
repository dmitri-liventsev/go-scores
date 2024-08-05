package vo

import (
	"database/sql/driver"
	"errors"
)

// TicketID ticket primary key.
type TicketID int

// Scan implementation of Scanner interface.
func (c *TicketID) Scan(value interface{}) error {
	if value == nil {
		*c = 0
		return nil
	}
	val, ok := value.(int64)
	if !ok {
		return errors.New("failed to scan CategoryID")
	}
	*c = TicketID(val)
	return nil
}

// Value implementation of Valuer interface.
func (c TicketID) Value() (driver.Value, error) {
	return int(c), nil
}

// ToInt converting value object to int value.
func (c TicketID) ToInt() int {
	return int(c)
}

// NewTicketID returns new TicketID instance.
func NewTicketID(id int) TicketID {
	return TicketID(id)
}
