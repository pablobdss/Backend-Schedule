# Backend Schedule - API em Golang

> 🇺🇸 [Read in English](./README.md)

API REST de agendamento para barbearias com autenticação via JWT, banco de dados PostgreSQL e estrutura modular. Projeto em construção por etapas. Os commits com atualizações estarei fazendo no Sábado ou Domingo, já que tenho mais tempo para organizar.

## Visão Geral

Esta API permite:

- Cadastro e login de usuários com senha criptografada
- Geração de JWT para autenticação
- Middleware de proteção de rotas
- Agendamentos (nas próximas versões)

Esta é a **Versão 1 (MVP Backend)**, focada em fundações sólidas:
estrutura de pastas, autenticação e integração com PostgreSQL.

---

## 🚦 Status da Versão

- ✅ Setup inicial do projeto
- ✅ Conexão com banco PostgreSQL via Docker
- ✅ Registro e login de usuários com validação
- ✅ Middleware de autenticação com JWT
- ✅ Rota protegida `/dashboard` testada
- 🔜 CRUD e Validações de agendamentos

---

## 📂 Estrutura de Pastas (V1)

- `cmd/api` # Ponto de entrada principal (main.go)
- `internal/db` # Conexão e utilitários do banco
- `internal/user` # Lógica de negócio, handlers, serviços
- `internal/auth` # Hash e Verify da senha | Geração e validação de JWT
- `internal/middleware` # Middlewares de autenticação, logger, etc
- `pkg/models` # Structs e DTOs compartilhados
- `docker-compose.yaml`  # Serviço PostgreSQL

---

## Estrutura do Projeto

## ⚙️ Tecnologias

- Go 1.23.4
- PostgreSQL 17
- Docker + Docker Compose
- JWT (via `github.com/golang-jwt/jwt/v5`)
- Variáveis de ambiente com `godotenv`

---

### Passos

```bash
1. Clone o repositório
  git clone https://github.com/pablobdss/Backend-Schedule.git
  cd Backend-Schedule
2. Crie o arquivo .env com suas variáveis:
  POSTGRES_USER=seu_user
  POSTGRES_PASSWORD=sua_senha
  POSTGRES_DB=seu_db
  POSTGRES_PORT=sua_porta
  JWT_SECRET=sua_senha_segura
3. Suba o banco de dados com Docker:
  docker-compose up -d
4. Instale as dependências:
  go mod tidy
5. Execute o projeto:
  go run cmd/api/main.go
```

---

## 📬 Endpoints (Versão 1)

### POST /register
```json
{
  "name": "João Barber",
  "email": "joao@example.com",
  "password": "secure123"
}
```
### POST /login
``` json
{
  "email": "seu@example.com",
  "password": "suapassword123"
}
```
### GET /dashboard
  Requer Authorization: Bearer <token>

---

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

---

## 📍 Roadmap de Evolução

### 🧱 Versão 2 – Infraestrutura & Segurança (em breve)

- [ ] Dockerizar o backend Go
- [ ] Middleware de autenticação com verificação de expiração do JWT
- [ ] Rate limiting básico (por IP)
- [ ] Middleware global de tratamento de erros
- [ ] Logs estruturados com zerolog ou zap
- [ ] Testes básicos nas rotas principais
- [ ] Atualização do README com uso via Docker

### ✨ Versão 3 – Experiência Real de Produto (visão futura)

- Integração com Frontend Typescript React/Next.js
- Chatbot explicativo com OpenAI
- Agendamentos com validação de horários
- Integração de pagamento (Simulado)

---

## Contribuindo

Este projeto faz parte do meu portfólio pessoal e meu aprendizado com Go. Feedbacks construtivos são muito bem vindos!
Sinta-se a vontade para abrir uma issue ou me chamar no LinkedIn!
