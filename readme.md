# Sobre

Uma aplicação web construída com Go, que implementa uma arquitetura monolítica modular, promovendo uma estrutura de código limpa e de fácil manutenção.

## Descrição

Este projeto demonstra a aplicação dos conceitos de uma arquitetura monolítica modular, onde as diferentes funcionalidades são organizadas em módulos distintos. A separação de responsabilidades facilita a manutenção, a escalabilidade e a legibilidade do código.


## Estrutura do Projeto

```
seu-projeto/
├── cmd/
│   └── app/
│       └── main.go            # Entry point da aplicação
├── docker-compose.yml
├── example.env
├── go.mod
├── go.sum
├── internal/
│   ├── core                   # Módulos de domínio
│   │   ├── books-management    # Módulo de gestão de livros
│   │   └── user-management      # Módulo de gestão de usuários
│   │       ├── domain/
│   │       │   └── user.go     # Pacote de domínio do módulo
│   │       ├── features/       # Funcionalidades do módulo
│   │       │   ├── create/
│   │       │   │   ├── data/
│   │       │   │   │   └── repository.go        # Persistência de dados
│   │       │   │   ├── http/
│   │       │   │   │   ├── handler.go            # Manipulação de rotas
│   │       │   │   │   └── routes.go
│   │       │   │   ├── use_case/
│   │       │   │   │   └── services.go           # Regras de negócio
│   │       │   │   └── web/
│   │       │   │       └── views/
│   │       │   │           ├── error.html
│   │       │   │           ├── register.html
│   │       │   │           └── success.html
│   │       │   ├── delete/
│   │       │   ├── find_all/
│   │       │   │   ├── data/
│   │       │   │   │   └── repository.go
│   │       │   │   ├── http/
│   │       │   │   │   ├── handler.go
│   │       │   │   │   └── routes.go
│   │       │   │   ├── use_case/
│   │       │   │   │   └── services.go
│   │       │   │   └── views/
│   │       │   │       └── list.html
│   │       │   └── update/
│   │       ├── module.go       # Carrega e configura o módulo
│   │       └── sql/
│   │           └── init.sql    # SQL entrypoint
│   └── shared/                # Módulos compartilhados
│       ├── config/
│       │   └── env_vars/
│       │       ├── load_database_vars/
│       │       │   └── execute.go
│       │       └── load_server_vars/
│       │           └── execute.go
│       └── database/
│           ├── configure_database/
│           │   └── execute.go
│           └── module.go
|
```
## Explicação
- **cmd/**: Contém o código de entrada da aplicação; o arquivo **main.go** é o ponto de partida da execução.

- **docker-compose.yml**: Configura e orquestra serviços para a aplicação em contêineres Docker.

- **internal/**: Lógica central da aplicação organizada em módulos.
  - **core/**: Módulos principais da lógica de domínio.
    - **books-management/**: Gestão de livros.
    - **user-management/**: Gestão de usuários.
      - **domain/**: Modelos e lógica de domínio de usuários.
      - **features/**: Implementações de funcionalidades do módulo.
        - **create/**: Funcionalidade para criar registros.
        - **delete/**: Funcionalidade para deletar registros.
        - **find_all/**: Funcionalidade para buscar todos os registros.
        - **update/**: Funcionalidade para atualizar registros.
      - **module.go**: Carrega e configura o módulo de usuários.
      - **sql/**: Scripts SQL para inicialização do banco de dados.

- **shared/**: Módulos utilizados em toda a aplicação.
  - **config/**: Carrega variáveis de ambiente para configuração.
  - **database/**: Configuração e inicialização do banco de dados.
