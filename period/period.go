package period

import "time"

func Create(start, end time.Time) Period {
	return Period{
		Start: start,
		End:   end,
	}
}

type Period struct {
	Start time.Time
	End   time.Time
}

func (p Period) OverlappingDays(target Period) int {
	start := p.Start
	if target.Start.After(p.Start) {
		start = target.Start
	}
	end := p.End
	if target.End.Before(p.End) {
		end = target.End
	}

	if end.Before(start) {
		return 0
	}
	return end.Day() - start.Day() + 1
}

func MonthDays(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
