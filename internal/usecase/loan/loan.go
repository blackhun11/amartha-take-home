package loan

import (
	"amartha-loan-system/internal/model"
	"amartha-loan-system/internal/repository/loan"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type loanUsecase struct {
	loan.PG
}

type LoanUsecase interface {
	CreateLoan(ctx context.Context, loan model.Loan) (model.Loan, error)
	ApproveLoan(ctx context.Context, loanApproval model.LoanApproval, loanID int64) (model.LoanApproval, error)
	RecordLoanInvestment(ctx context.Context, loanInvestment model.LoanInvestment, loanID int64) (model.LoanInvestment, error)
	RecordDisburseLoan(ctx context.Context, loanDisbursement model.LoanDisbursement, loanID int64) (model.LoanDisbursement, error)
}

func NewloanUsecase(pg loan.PG) LoanUsecase {
	return &loanUsecase{
		PG: pg,
	}
}

func (u *loanUsecase) CreateLoan(ctx context.Context, loan model.Loan) (model.Loan, error) {
	loanID, err := u.PG.CreateLoan(ctx, loan)
	if err != nil {
		return model.Loan{}, err
	}

	// asuming this is the process of generating agreement letter link
	loan.AgreementLetterLink = uuid.NewString()
	go func() {
		_ = u.PG.InsertTransitionLog(ctx, loanID, "", model.StateProposed)
	}()

	loan.LoanID = loanID

	return loan, err
}

func (u *loanUsecase) ApproveLoan(ctx context.Context, loanApproval model.LoanApproval, loanID int64) (model.LoanApproval, error) {
	// TODO: use transaction later

	err := u.PG.RecordLoanApproval(ctx, loanID, loanApproval)
	if err != nil {
		return model.LoanApproval{}, err
	}

	err = u.PG.UpdateLoanState(ctx, loanID, model.StateApproved)
	if err != nil {
		return model.LoanApproval{}, err
	}

	go func() {
		_ = u.PG.InsertTransitionLog(ctx, loanID, model.StateProposed, model.StateApproved)
	}()

	loanApproval.LoanID = loanID

	return loanApproval, nil
}

func (u *loanUsecase) RecordLoanInvestment(ctx context.Context, loanInvestment model.LoanInvestment, loanID int64) (model.LoanInvestment, error) {
	loan, err := u.PG.GetLoanByID(ctx, loanID)
	if err != nil {
		return model.LoanInvestment{}, err
	}

	if loan.State != model.StateApproved {
		return model.LoanInvestment{}, fmt.Errorf("loan must be in approved state to accept investments")
	}

	currentAmount, err := u.PG.GetLoanInvestmentAmountCount(ctx, loanID)
	if err != nil {
		return model.LoanInvestment{}, err
	}
	totalInvested := loanInvestment.Amount + currentAmount

	if totalInvested > loan.PrincipalAmount {
		return model.LoanInvestment{}, fmt.Errorf("total investment cannot exceed principal amount")
	}

	// TODO: use transaction later
	err = u.PG.RecordLoanInvestment(ctx, loanID, loanInvestment)
	if err != nil {
		return model.LoanInvestment{}, err
	}

	if totalInvested == loan.PrincipalAmount {
		err = u.PG.UpdateLoanState(ctx, loanID, model.StateInvested)
		if err != nil {
			return model.LoanInvestment{}, err
		}

		go func() {
			// asumming this part was the processof sending an email containing link to agreement letter via MQ
		}()
	}

	go func() {
		_ = u.PG.InsertTransitionLog(ctx, loanID, model.StateApproved, model.StateInvested)
	}()

	return loanInvestment, nil
}

func (u *loanUsecase) RecordDisburseLoan(ctx context.Context, loanDisbursement model.LoanDisbursement, loanID int64) (model.LoanDisbursement, error) {
	err := u.PG.RecordLoanDisbursement(ctx, loanID, loanDisbursement)
	if err != nil {
		return model.LoanDisbursement{}, err
	}

	go func() {
		_ = u.PG.InsertTransitionLog(ctx, loanID, model.StateInvested, model.StateDisbursed)
	}()

	loanDisbursement.LoanID = loanID

	return loanDisbursement, nil
}
