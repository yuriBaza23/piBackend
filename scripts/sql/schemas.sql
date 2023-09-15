create table if not exists users (
  id uuid default gen_random_uuid() primary key, 
  name varchar not null, 
  email varchar not null,
  password varchar not null,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);

create table if not exists users_companies (
  id uuid default gen_random_uuid() primary key, 
  companyId varchar default '',
  userId uuid not null,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);
