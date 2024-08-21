package loan

import (
	"amartha-loan-system/internal/model"
	"context"
)

type Transition interface {
	InsertTransitionLog(ctx context.Context, loanID int64, fromState model.State, toState model.State) error
}

func (r *pgRepository) InsertTransitionLog(ctx context.Context, loanID int64, fromState model.State, toState model.State) error {
	query := `
		INSERT INTO loan_state_transitions (loan_id, from_state, to_state)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.MasterConn.ExecContext(ctx, query, loanID, fromState, toState)
	if err != nil {
		return err
	}

	return nil
}
