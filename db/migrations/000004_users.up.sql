CREATE TABLE IF NOT EXISTS public.users (
  id int8 NOT NULL,
  role_id int8 NULL REFERENCES roles(id) ON DELETE CASCADE,
  first_name varchar NULL,
  second_name varchar NULL,
  email varchar NULL,
  phone_number varchar NULL,
  adress varchar NULL,
  birth_date date NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id),
  CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES public.roles (id)
);