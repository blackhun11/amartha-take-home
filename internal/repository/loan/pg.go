package loan

import (
	"amartha-loan-system/internal/pg"
)

type PG interface {
	Loan
	Transition
	LoanApproval
	LoanInvestment
	LoanDisbursement
}

type pgRepository struct {
	db pg.DB
}

func NewPG(db pg.DB) PG {
	return &pgRepository{
		db: db,
	}
}
