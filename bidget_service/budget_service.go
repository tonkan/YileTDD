package bidget_service

import (
	"YileTDD/budget_repo"
	"time"
)

type BudgetService struct {
	br budget_repo.IBudgetRepo
}

func (bs BudgetService) Query(start, end time.Time) (amount float64) {
	if end.Before(start) {
		return
	}
	for _, b := range bs.br.GetAll() {
		amount += float64(b.AmountInInterval(start, end))
	}
	return
}
