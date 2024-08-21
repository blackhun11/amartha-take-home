package request

import (
	"amartha-loan-system/internal/model"
	"time"
)

type LoanApprovalRequest struct {
	FieldValidatorEmployeeID int64     `json:"field_validator_employee_id" validate:"required"`
	ProofOfVisit             string    `json:"proof_of_visit" validate:"required"`
	ApprovalDate             time.Time `json:"approval_date" validate:"required"`
}

func (r LoanApprovalRequest) ToModel() model.LoanApproval {
	return model.LoanApproval{
		FieldValidatorEmployeeID: r.FieldValidatorEmployeeID,
		ProofOfVisit:             r.ProofOfVisit,
		ApprovalDate:             r.ApprovalDate,
	}
}
