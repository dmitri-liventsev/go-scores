package vo

// PeriodID period primary key.
type PeriodID int

// ToInt convert value object to int value.
func (p PeriodID) ToInt() int {
	return int(p)
}

// NewPeriodID returns new PeriodID instance.
func NewPeriodID(id int) PeriodID {
	return PeriodID(id)
}
