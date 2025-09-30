# Receitas App

Este projeto é uma API REST desenvolvida em Golang utilizando o framework Fiber, com autenticação JWT e integração com banco de dados PostgreSQL. O objetivo é gerenciar usuários e receitas, permitindo cadastro, login, listagem de usuários e receitas, além de criação de receitas apenas por administradores.

## Funcionalidades

- Cadastro de usuários com hash de senha (bcrypt)
- Login com geração de token JWT
- Listagem de usuários e receitas
- Criação de receitas (restrito a administradores)
- Proteção de rotas via middleware JWT

## Estrutura do Projeto
```
receitas_app/
├── backend/
│   ├── cmd/
│   │   └── main.go
│   ├── config/
│   │   └── db.go
│   ├── controllers/
│   │   ├── auth_controllers.go
│   │   └── receitas_controllers.go
│   ├── middlewares/
│   │   └── auth.go
│   ├── models/
│   │   ├── receitas.go
│   │   └── user.go
│   ├── routes/
│   │   └── routes.go
│   └── utils/
│       ├── hash.go
│       └── jwt.go
├── .env
├── .gitignore
├── compose.yml
├── go.mod
└── go.sum
```


## Como executar

1. Clone o repositório
```
git clone https://github.com/Whofelisberto/receitas_app.git
```
3. Configure o arquivo `.env` com a variável `JWT_SECRET`
4. Suba o banco de dados com Docker Compose:
   
   ```sh
   docker compose up -d
   ```

## Instale as dependências Go:
    ```go mod tidy ```


## Execute a aplicação:
 ```go run backend/cmd/main.go```


 # Endpoints

 ``` 
POST /registrar — Cadastro de usuário
POST /login — Login e obtenção de token JWT
GET /users — Listagem de usuários
GET /users/:id — Buscar usuário por ID
POST /receitas — Criar receita (apenas admin)
GET /receitas — Listar receitas
 ```

## Tecnologias Usadas:
 ```
Go
Fiber
GORM
PostgreSQL
JWT
Docker Compose
 ```


## Observações:
```
Apenas usuários com papel admin podem criar receitas.
As senhas são armazenadas de forma segura utilizando bcrypt.
O projeto está pronto para ser expandido com novas funcionalidades.
