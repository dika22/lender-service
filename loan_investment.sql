-- public.loan_investments definition

-- Drop table

-- DROP TABLE public.loan_investments;

CREATE TABLE public.loan_investments (
	id uuid NOT NULL,
	loan_id varchar(50) NOT NULL,
	investor_id varchar(50) NOT NULL,
	amount int8 NOT NULL,
	created_at timestamp NULL DEFAULT now(),
	CONSTRAINT loan_investments_pkey PRIMARY KEY (id)
);


-- public.loan_investments foreign keys

ALTER TABLE public.loan_investments ADD CONSTRAINT loan_investments_loan_id_fkey FOREIGN KEY (loan_id) REFERENCES public.loans(id);