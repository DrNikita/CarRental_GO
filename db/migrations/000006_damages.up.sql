CREATE TABLE IF NOT EXISTS public.damages (
  id int8 NOT NULL,
  user_id int8 NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  car_id int8 NOT NULL REFERENCES cars(id) ON DELETE CASCADE,
  description varchar NULL,
  price int8 NOT NULL,
  is_paid boolean NOT NULL,
  CONSTRAINT damages_pkey PRIMARY KEY (id),
  CONSTRAINT damages_user_fk FOREIGN KEY (user_id) REFERENCES public.users (id),
  CONSTRAINT damages_car_fk FOREIGN KEY (car_id) REFERENCES public.cars (id)
);