-- Uma incubadora possui um ID, nome, senha e email. Seguindo o padrão, possuem datas de criação e atualização
-- Empresas incubadas devem possuir uma foreign key que aponta para o id da incubadora da qual faz parte

create table if not exists incubators (
    id uuid default gen_random_uuid() primary key, 
    name varchar not null,
    password varchar not null,
    email varchar not null,
    createdAt timestamp default now(),
    updatedAt timestamp default now()
);

