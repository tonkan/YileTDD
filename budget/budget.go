package budget

import (
	"YileTDD/period"
	"time"
)

func Create(yearMonth string, amount int) Budget {
	return Budget{
		yearMonth: yearMonth,
		amount:    amount,
	}
}

type Budget struct {
	yearMonth string
	amount    int
}

func (b Budget) YearMonth() string {
	return b.yearMonth
}

func (b Budget) Amount() int {
	return b.amount
}

func (b Budget) OverlappingAmount(p period.Period) int {
	days := p.OverlappingDays(period.Create(b.firstDay(), b.lastDay()))
	return b.dailyAmount() * days
}

func (b Budget) firstDay() time.Time {
	t, _ := time.Parse("200601", b.YearMonth())
	return t
}

func (b Budget) lastDay() time.Time {
	t, _ := time.Parse("200601", b.YearMonth())
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
}

func (b Budget) dailyAmount() int {
	t, _ := time.Parse("200601", b.yearMonth)
	return b.Amount() / period.MonthDays(t)
}
