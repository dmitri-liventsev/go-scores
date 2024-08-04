package vo

type PeriodID int

func (p PeriodID) ToInt() int {
	return int(p)
}

func NewPeriodID(id int) PeriodID {
	return PeriodID(id)
}
