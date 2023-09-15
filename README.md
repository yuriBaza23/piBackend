<br/>

<!-- <p align="center"><a href="https://newhappen.com.br" target="_blank"><img src="https://github.com/NewHappen-Company/oficial/blob/master/src/assets/logoBlue.svg?raw=true" height="70"></a></p> -->

<p align="center">
<h1 align="center">Projeto Integrador</h1>
</p>

<br/>

<p align="center">
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Lang" />
    <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white" alt="Postgres" />
    <img src="https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
</p>

## Sobre
Uma aplicação para gestão de empresas por uma incubadora ou aceleradora, onde o responsável por uma incubadora ou aceleradora pode gerenciar as empresas que estão incubadas ou aceleradas por ela, e as empresas podem gerenciar seus profissionais e seus projetos.

O objetivo é que a empresa consiga, em uma única plataforma, gerenciar seus colaboradores, compromissos, contratos, receitas e despesas. Em contrapartida, a incubadora ou aceleradora consegue gerenciar as empresas que estão incubadas ou aceleradas por ela de tal forma que consiga ter uma visão geral de todas as empresas e de cada uma individualmente.

A aplicação contará com um chat entre as duas partes, onde a incubadora ou aceleradora pode se comunicar com as empresas e vice-versa, além das funcionalidades já mencionadas anteriormente.

## :warning: Arquivos importantes
config.toml
-------------

O arquivo ***config.toml*** deve ser modificado com base nas configurações do seu banco de dados e porta da api.
  
Esse arquivo deve estar na pasta raiz do projeto e conter obrigatóriamente as seguintes informações:
```toml
[api]
port=3333

[db]
host=""
port=""
user=""
password=""
database=""
```
>EM HIPÓTESE ALGUMA SUBA SEU ARQUIVO DE CONFIGURAÇÃO PARA O GITHUB
  
Caso seja necessário criar uma nova variável de desenvolvimento, observe que há um padrão:
```toml
[ESCOPO]
NOME_DA_VARIAVEL = "valor"
```

_Caso tenha criado uma variável, especifique isso em um Pull Request_  

-----------------

## Requisitos
- [x] Configurar config.toml
- [x] Instalar dependências com `go mod tidy`

## Rodando o backend
Execute o backend com
```shell
go run ./cmd/pi/main.go
```

## Criando seu banco de dados
Para criar seu banco de dados, execute o seguinte comando:
```shell
docker run -d --name api-pi -p 5433:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5
```
O comando acima irá criar um container com o banco de dados postgresql na porta 5433. Para acessar o banco de dados, utilize o seguinte comando:
```shell
docker exec -it api-pi psql -U postgres
```
Agora, para criar o banco de dados, execute o seguinte comando:
```sql
create database pi;
```
Após a criação do banco de dados, você poderá criar um usuário para acessar o bd. Caso não queira criar um usuário, deverá utilizar o usuário padrão do postgres no seu arquivo de configuração. Para criar um usuário, execute o seguinte comando:
```sql
create user user_pi;
```
Para definir uma senha para o usuário:
```sql
alter user user_pi with encrypted password '1234';
```
Para dar permissões ao usuário:
```sql
grant all privileges on database pi to user_pi;
grant all privileges on all tables in schema public to user_pi;
grant all privileges on all sequences in schema public to user_pi;
```
Para se conectar ao banco de dados com o usuário criado, execute o seguinte comando:
```sql
\c pi;
```
Por fim, criaremos nossas tabelas com base no arquivo `schema.sql`.

Para sair do banco de dados, execute o comando `exit` no seu terminal.
Lembre-se de configurar o arquivo de configuração com as informações do seu banco de dados. 