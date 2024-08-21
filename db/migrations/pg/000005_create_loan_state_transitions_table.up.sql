BEGIN;

CREATE TABLE loan_state_transitions (
    id SERIAL PRIMARY KEY,
    loan_id INT REFERENCES loans(id) ON DELETE CASCADE,
    from_state TEXT NOT NULL,
    to_state TEXT NOT NULL,
    transitioned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


COMMIT;
