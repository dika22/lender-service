-- public.loans definition

-- Drop table

-- DROP TABLE public.loans;

CREATE TABLE public.loans (
	id varchar NOT NULL,
	borrower_id varchar(50) NOT NULL,
	employee_id varchar(50),
	principal_amount int8 NOT NULL,
	rate numeric(5, 4) NOT NULL,
	roi numeric(5, 4) NOT NULL,
	state int4 NOT NULL DEFAULT 1,
	agreement_url text NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL DEFAULT now(),
	CONSTRAINT loans_pkey PRIMARY KEY (id)
);