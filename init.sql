DROP TABLE IF EXISTS deduction CASCADE;
CREATE TABLE public.deduction (
	deduction_id int4 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1 NO CYCLE) NOT NULL,
	deduction_type varchar NOT NULL,
	deduction_amount numeric NOT NULL,
	CONSTRAINT decuction_pk PRIMARY KEY (deduction_id)
);
CREATE UNIQUE INDEX deduction_deduction_type_idx ON public.deduction USING btree (deduction_type);

INSERT INTO deduction (deduction_type, deduction_amount) VALUES
('k-receipt', 50000),
('personal_deduction', 60000),
('donation', 100000);
