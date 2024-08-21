CREATE TABLE loan_disbursements (
    id SERIAL PRIMARY KEY,
    loan_id INT REFERENCES loans(id) ON DELETE CASCADE,
    agreement_letter_signed TEXT NOT NULL,  -- URL or path to the signed agreement (PDF/JPEG)
    field_officer_employee_id TEXT NOT NULL,
    disbursement_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
