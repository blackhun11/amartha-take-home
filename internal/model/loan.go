package model

import "time"

type State string

const (
	StateProposed  State = "PROPOSED"
	StateApproved  State = "APPROVED"
	StateInvested  State = "INVESTED"
	StateDisbursed State = "DISBURSED"
)

type Loan struct {
	LoanID              int64   `json:"loan_id"`
	BorrowerID          int64   `json:"borrower_id"`
	PrincipalAmount     float64 `json:"principal_amount"`
	AgreementLetterLink string  `json:"agreement_letter_link"`
	Rate                float64 `json:"rate"`
	ROI                 float64 `json:"roi"`
	State               State   `json:"state"`
}

type LoanApproval struct {
	LoanID                   int64     `json:"loan_id"`
	FieldValidatorEmployeeID int64     `json:"field_validator_employee_id"`
	ProofOfVisit             string    `json:"proof_of_visit"`
	ApprovalDate             time.Time `json:"approval_date"`
}

type LoanInvestment struct {
	LoanID       int64   `json:"loan_id"`
	InvestmentID int64   `json:"investment_id"`
	Amount       float64 `json:"amount"`
}

type LoanDisbursement struct {
	LoanID                 int64     `json:"loan_id"`
	AggrementLetterSigned  string    `json:"agreement_letter_signed"`
	FieldOfficerEmployeeID int64     `json:"field_officer_employee_id"`
	DisbursementDate       time.Time `json:"disbursement_date"`
}
