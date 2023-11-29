create table if not exists projects (
  id uuid default gen_random_uuid() primary key, 
  name varchar not null,
  description varchar,
  companyID uuid not null,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);

create table if not exists warnings (
  id uuid default gen_random_uuid() primary key, 
  title varchar not null,
  content varchar not null,
  cmpID varchar not null default '',
  incID varchar not null default '',
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);

create table if not exists incubators (
  id uuid default gen_random_uuid() primary key, 
  name varchar not null,
  email varchar not null,
  password varchar not null,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);

create table if not exists users (
  id uuid default gen_random_uuid() primary key, 
  name varchar not null, 
  email varchar not null,
  password varchar not null,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);

alter table users add column if not exists isPreRegister boolean default false;

create table if not exists users_companies (
  id uuid default gen_random_uuid() primary key, 
  companyId varchar default '',
  userId uuid not null,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);

alter table users_companies add column if not exists type varchar default 'other';

create table if not exists companies (
  id uuid default gen_random_uuid() primary key, 
  name varchar not null,
  email varchar not null,
  cnpj varchar not null,
  hubId varchar not null default '',
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);

create table if not exists finances (
  id uuid default gen_random_uuid() primary key, 
  name varchar not null,
  type varchar not null,
  value integer not null default 0,
  companyId uuid not null,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);

create table if not exists tasks (
  id uuid default gen_random_uuid() primary key, 
  title varchar not null,
  description varchar,
  status varchar not null,
  initialDate timestamp,
  finalDate timestamp,
  projectId uuid not null,
  companyId uuid not null,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);

create table if not exists categories(
  id uuid default gen_random_uuid() primary key, 
  name varchar not null,
  companyId uuid not null,
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);
