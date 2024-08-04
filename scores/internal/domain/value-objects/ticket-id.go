package vo

import (
	"database/sql/driver"
	"errors"
)

type TicketID int

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

func (c TicketID) Value() (driver.Value, error) {
	return int(c), nil
}

func NewTicketID(id int) TicketID {
	return TicketID(id)
}

func (c TicketID) ToInt() int {
	return int(c)
}
