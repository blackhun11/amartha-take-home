BEGIN;

    CREATE TABLE loans (
        id SERIAL PRIMARY KEY,
        borrower_id INT NOT NULL,
        principal_amount DOUBLE PRECISION NOT NULL,
        rate DOUBLE PRECISION NOT NULL,
        roi DOUBLE PRECISION NOT NULL,
        state TEXT DEFAULT 'PROPOSED',
        link_to_agreement_letter TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );

COMMIT;