package budget

import "time"

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

func (b Budget) DailyAmount() int {
	t, _ := time.Parse("200601", b.yearMonth)
	return b.Amount() / b.daysIn(t.Month(), t.Year())
}

func (b Budget) AmountInInterval(start, end time.Time) int {
	t, _ := time.Parse("200601", b.yearMonth)

	// 起訖同年月
	if b.sameYearMonth(t, start) && b.sameYearMonth(t, end) {
		return b.DailyAmount() * (end.Day() - start.Day() + 1)
	}
	// 開始時間撞柱
	if b.sameYearMonth(t, start) {
		monthDays := b.daysIn(start.Month(), start.Year())
		return (monthDays - start.Day() + 1) * b.DailyAmount()
	}
	// 結束時間撞柱
	if b.sameYearMonth(t, end) {
		return end.Day() * b.DailyAmount()
	}
	// 指定區間內的完整月份
	if start.Before(t) && end.After(t) {
		return b.amount
	}
	return 0
}

func (b Budget) sameYearMonth(source time.Time, target time.Time) bool {
	return target.Year() == source.Year() && target.Month() == source.Month()
}

func (b Budget) daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
