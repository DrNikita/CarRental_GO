CREATE TABLE IF NOT EXISTS public.orders (
  id int8 NOT NULL,
  user_id int8 NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  car_id int8 NOT NULL REFERENCES cars(id) ON DELETE CASCADE,
  start_date date NOT NULL,
  end_date date NOT NULL,
  price int8 NOT NULL,
  rejectionReason varchar NULL,
  status boolean NOT NULL,
  is_paid boolean NOT NULL,
  CONSTRAINT orders_pkey PRIMARY KEY (id),
  CONSTRAINT orders_user_fk FOREIGN KEY (user_id) REFERENCES public.users (id),
  CONSTRAINT orders_car_fk FOREIGN KEY (car_id) REFERENCES public.cars (id)
);