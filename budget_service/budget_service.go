package budget_service

import (
	"YileTDD/budget_repo"
	"YileTDD/period"
	"time"
)

type BudgetService struct {
	br budget_repo.IBudgetRepo
}

func (bs BudgetService) Query(start, end time.Time) (amount float64) {
	if end.Before(start) {
		return
	}
	p := period.Create(start, end)
	for _, b := range bs.br.GetAll() {
		amount += float64(b.OverlappingAmount(p))
	}
	return
}
