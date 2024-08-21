package loan

import (
	"amartha-loan-system/internal/model"
	"context"
)

type LoanDisbursement interface {
	RecordLoanDisbursement(ctx context.Context, loanID int64, loanDisbursement model.LoanDisbursement) error
}

func (r *pgRepository) RecordLoanDisbursement(ctx context.Context, loanID int64, loanDisbursement model.LoanDisbursement) error {
	query := `
		INSERT INTO loan_investments (loan_id, agreement_letter_signed, field_officer_employee_id, disbursement_date)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.MasterConn.ExecContext(ctx, query, loanID, loanDisbursement.AggrementLetterSigned, loanDisbursement.FieldOfficerEmployeeID, loanDisbursement.DisbursementDate)
	if err != nil {
		return err
	}

	return nil
}
