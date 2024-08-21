package request

import "amartha-loan-system/internal/model"

type CreateLoanRequest struct {
	BorrowerID      int64   `json:"borrower_id" validate:"required"`
	PrincipalAmount float64 `json:"principal_amount" validate:"required"`
	Rate            float64 `json:"rate" validate:"required"`
	ROI             float64 `json:"roi" validate:"required"`
}

func (r CreateLoanRequest) ToModel() model.Loan {
	return model.Loan{
		BorrowerID:      r.BorrowerID,
		PrincipalAmount: r.PrincipalAmount,
		Rate:            r.Rate,
		ROI:             r.ROI,
		State:           model.StateProposed,
	}
}
