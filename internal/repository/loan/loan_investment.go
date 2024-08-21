package loan

import (
	"amartha-loan-system/internal/model"
	"context"
)

type LoanInvestment interface {
	RecordLoanInvestment(ctx context.Context, loanID int64, loanInvestment model.LoanInvestment) error
	GetLoanInvestmentAmountCount(ctx context.Context, loanID int64) (float64, error)
}

func (r *pgRepository) RecordLoanInvestment(ctx context.Context, loanID int64, loanInvestment model.LoanInvestment) error {
	query := `
		INSERT INTO loan_investments (loan_id, investment_id, amount)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.MasterConn.ExecContext(ctx, query, loanID, loanInvestment.InvestmentID, loanInvestment.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) GetLoanInvestmentAmountCount(ctx context.Context, loanID int64) (float64, error) {
	query := `
		SELECT COUNT(amount)
		FROM loan_investments
		WHERE loan_id = $1;
	`

	var amount float64
	err := r.db.ReplicaConn.QueryRowContext(ctx, query, loanID).Scan(
		&amount,
	)
	if err != nil {
		return amount, err
	}

	return amount, nil
}
