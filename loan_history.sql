-- Drop table if exists (optional)
-- DROP TABLE IF EXISTS public.loan_history;

CREATE TABLE public.loan_history (
	id SERIAL PRIMARY KEY,
	loan_id varchar NOT NULL,
	employee_id varchar(50) NOT NULL,
	state integer DEFAULT 1, 
	comments TEXT,
	approval_picture_url TEXT,
	signed_agreement_url TEXT, 
	approval_date varchar,
	disbursement_date varchar,
	created_at timestamp NULL DEFAULT now()
);

-- Add foreign key constraint properly
ALTER TABLE public.loan_history
ADD CONSTRAINT loan_history_loan_id_fkey
FOREIGN KEY (loan_id)
REFERENCES public.loans(id);
