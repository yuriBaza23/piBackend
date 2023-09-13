create table if not exists incubators (
  id uuid default gen_random_uuid() primary key, 
  name varchar not null,
  email varchar not null,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);
