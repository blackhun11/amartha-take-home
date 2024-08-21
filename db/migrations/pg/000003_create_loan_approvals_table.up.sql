CREATE TABLE loan_approvals (
    id SERIAL PRIMARY KEY,
    loan_id INT REFERENCES loans(id) ON DELETE CASCADE,
    field_validator_employee_id INT NOT NULL,
    proof_of_visit TEXT NOT NULL,  -- URL or path to the image/file
    approval_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
