CREATE TABLE IF NOT EXISTS public.roles (
  id int8 NOT NULL,
  role_name varchar NULL,
  CONSTRAINT roles_pkey PRIMARY KEY (id)
);

INSERT INTO public.roles (id, role_name) VALUES (1, 'user');

INSERT INTO public.roles (id, role_name) VALUES (2, 'admin');