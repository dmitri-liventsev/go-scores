package entities

import (
	vo "go-scores/scores/internal/domain/value-objects"
	pkgtime "go-scores/scores/pkg/time"
	"time"
)

// Period period O_o.
type Period struct {
	ID    vo.PeriodID
	Start time.Time
	End   time.Time
	Type  string
}

// GetInvariant get period invariant.
func (p Period) GetInvariant() string {
	return pkgtime.GetPeriodInvariant(p.Start, p.Type)
}

// IsDateBetween returns true if date are between start and end of period.
func (p Period) IsDateBetween(date time.Time) bool {
	return pkgtime.IsDateBetween(date, p.Start, p.End)
}

// NewPeriod returns new period instance.
func NewPeriod(id int, start, end time.Time, typ string) Period {
	return Period{
		ID:    vo.NewPeriodID(id),
		Start: start,
		End:   end,
		Type:  typ,
	}
}
