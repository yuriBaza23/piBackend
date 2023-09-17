create table if not exists warnings (
  id uuid default gen_random_uuid() primary key, 
  title varchar not null,
  content varchar not null,
  cmpID varchar not null default '',
  incID varchar not null default '',
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);
