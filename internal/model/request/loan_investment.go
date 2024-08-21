package request

import "amartha-loan-system/internal/model"

type LoanInvestmentRequest struct {
	InvestmentID int64   `json:"investment_id" validate:"required"`
	Amount       float64 `json:"amount" validate:"required"`
}

func (r LoanInvestmentRequest) ToModel() model.LoanInvestment {
	return model.LoanInvestment{
		InvestmentID: r.InvestmentID,
		Amount:       r.Amount,
	}
}
