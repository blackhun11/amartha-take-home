package loan

import (
	"amartha-loan-system/internal/model"
	"context"
)

type LoanApproval interface {
	RecordLoanApproval(ctx context.Context, loanID int64, loanApproval model.LoanApproval) error
}

func (r *pgRepository) RecordLoanApproval(ctx context.Context, loanID int64, loanApproval model.LoanApproval) error {
	query := `
		INSERT INTO loan_approvals (loan_id, field_validator_employee_id, proof_of_visit, approval_date)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.MasterConn.ExecContext(ctx, query, loanID, loanApproval.FieldValidatorEmployeeID, loanApproval.ProofOfVisit, loanApproval.ApprovalDate)
	if err != nil {
		return err
	}

	return nil
}
