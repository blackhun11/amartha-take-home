package request

import (
	"amartha-loan-system/internal/model"
	"time"
)

type LoanDisbursementRequest struct {
	AggrementLetterSigned  string    `json:"agreement_letter_signed" validate:"required"`
	FieldOfficerEmployeeID int64     `json:"field_officer_employee_id" validate:"required"`
	DisbursementDate       time.Time `json:"disbursement_date" validate:"required"`
}

func (r LoanDisbursementRequest) ToModel() model.LoanDisbursement {
	return model.LoanDisbursement{
		AggrementLetterSigned:  r.AggrementLetterSigned,
		FieldOfficerEmployeeID: r.FieldOfficerEmployeeID,
		DisbursementDate:       r.DisbursementDate,
	}
}
