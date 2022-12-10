package budget_repo

import "YileTDD/budget"

type IBudgetRepo interface {
	GetAll() []budget.Budget
}
