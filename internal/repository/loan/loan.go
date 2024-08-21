package loan

import (
	"amartha-loan-system/internal/model"
	"context"
)

type Loan interface {
	CreateLoan(ctx context.Context, loan model.Loan) (int64, error)
	GetLoanByID(ctx context.Context, loanID int64) (model.Loan, error)
	UpdateLoan(ctx context.Context, loan model.Loan) error
	DeleteLoan(ctx context.Context, loanID int64) error
	UpdateLoanState(ctx context.Context, loanID int64, state model.State) error
}

func (r *pgRepository) CreateLoan(ctx context.Context, loan model.Loan) (int64, error) {
	var loanID int64
	query := `
		INSERT INTO loans (borrower_id, principal_amount, rate, roi, state)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING loan_id;
	`

	err := r.db.MasterConn.QueryRowContext(ctx, query, loan.BorrowerID, loan.PrincipalAmount, loan.Rate, loan.ROI, loan.State).Scan(&loanID)
	if err != nil {
		return 0, err
	}

	return loanID, nil
}

func (r *pgRepository) GetLoanByID(ctx context.Context, loanID int64) (model.Loan, error) {
	query := `
		SELECT loan_id, borrower_id, principal_amount, rate, roi, state
		FROM loans
		WHERE loan_id = $1;
	`

	var loan model.Loan
	err := r.db.ReplicaConn.QueryRowContext(ctx, query, loanID).Scan(
		&loan.LoanID,
		&loan.BorrowerID,
		&loan.PrincipalAmount,
		&loan.Rate,
		&loan.ROI,
		&loan.State,
	)
	if err != nil {
		return model.Loan{}, err
	}

	return loan, nil
}

func (r *pgRepository) UpdateLoan(ctx context.Context, loan model.Loan) error {
	query := `
		UPDATE loans
		SET borrower_id = $1, principal_amount = $2, rate = $3, roi = $4, state = $5
		WHERE loan_id = $6;
	`

	_, err := r.db.MasterConn.ExecContext(ctx, query, loan.BorrowerID, loan.PrincipalAmount, loan.Rate, loan.ROI, loan.State, loan.LoanID)
	if err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) DeleteLoan(ctx context.Context, loanID int64) error {
	query := `
		DELETE FROM loans
		WHERE loan_id = $1;
	`

	_, err := r.db.MasterConn.ExecContext(ctx, query, loanID)
	if err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) UpdateLoanState(ctx context.Context, loanID int64, state model.State) error {
	query := `
		UPDATE loans
		SET state = $1
		WHERE loan_id = $2;
	`

	_, err := r.db.MasterConn.ExecContext(ctx, query, state, loanID)
	if err != nil {
		return err
	}

	return nil
}
