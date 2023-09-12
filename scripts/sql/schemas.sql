create database pi;

create user user_pi;

alter user user_pi with encrypted password '1234';

grant all privileges on database pi to user_pi;
grant all privileges on all tables in schema public to user_pi;
grant all privileges on all sequences in schema public to user_pi;