CREATE TABLE IF NOT EXISTS public.cars (
  id int8 NOT NULL,
  govnum varchar NULL,
  brand varchar NULL,
  issue_date date NULL,
  car_cost int8 NULL,
  rental_cost int8 NULL,
  CONSTRAINT cars_pkey PRIMARY KEY (id)
);