create table if not exists partners (
  id uuid default gen_random_uuid() primary key, 
  name varchar not null, 
  type varchar not null, 
  email varchar not null,
  companyId varchar not null default '',
  accountId varchar,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);