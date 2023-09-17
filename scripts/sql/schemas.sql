create table if not exists companies (
  id uuid default gen_random_uuid() primary key, 
  name varchar not null,
  companyId varchar not null,
  description varchar not null default '',
  createdAt timestamp default now(),
  updatedAt timestamp default now()s
);

