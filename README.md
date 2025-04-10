# Backend de Agendamentos - Golang

## Visão Geral

Projeto em Desenvolvimento de uma API para uma barbearia com horarios de agendamento usando autenticação via JWT e persistencia em PostgreSQL
Essa é a versao V1, com foco apenas na estrutura e MVP do projeto.

## Status da Versão

Etapa 1 concluída - Setup inicial do projeto, conexão com banco de dados e estrutura configurada.

Próxima etapa: Criar as rotas '/register' e '/login'

## 📂 Estrutura de Pastas (V1)

- `cmd/api` – Ponto de entrada da aplicação
- `internal/db` – Conexão com o banco de dados
- `internal/user` – Registro, login e lógica de usuário
- `pkg/models` – Structs compartilhadas

## Estrutura do Projeto

## ⚙️ Tecnologias

- Go 1.23.4
- PostgreSQL
- Docker
- JWT (futuramente)
- Dotenv (para variáveis de ambiente)

### Passos

```bash
git clone https://github.com/pablobdss/Backend-Schedule.git
cd Backend-Schedule
docker-compose up -d
go mod tidy
```

## 🐳 Docker

O banco de dados PostgreSQL pode ser executado via Docker para facilitar o setup local.
```
services:
  postgres:
    image: postgres:17
    container_name: schedule-postgres
    restart: always
    environment:
      POSTGRES_DB: ${POSTGRES_DB}
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
    ports:
      - "${POSTGRES_PORT}:${POSTGRES_PORT}"
```
