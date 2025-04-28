# Backend Schedule - Golang API

> 🇧🇷 [Leia em Português](./README.pt-br.md)

REST API for barbershop appointments using JWT authentication, PostgreSQL database, and modular project structure. Built in stages for professional portfolio purposes. Commits with Update will be on Saturdays or Sundays, since i have more free time to prepare.

## Overview

This API provides:

- User registration and login with encrypted passwords
- JWT token generation for authentication
- Route protection via middleware
- Appointments

You're viewing **Version 1 (Backend MVP)** — focused on solid foundations:
folder structure, authentication, and PostgreSQL integration.

---

## 🚦 Version Status

- ✅ Initial project setup
- ✅ PostgreSQL database connection via Docker
- ✅ User registration and login with validation
- ✅ JWT-based authentication middleware
- ✅ Protected `/dashboard` route tested
- ✅ Appointment CRUD (using Chi router for HTTP method matching)
- 🔄 Validations (no overlapping slots, past‐date bookings, enforce 08:00‑18:00) & structured error logging

---

## 📂 Folder Structure (V1)

- `cmd/api` – Main entry point (main.go)
- `internal/db` – Database connection and utilities
- `internal/user` – Business logic, handlers, and services
- `internal/auth` – Password hash/verify | JWT generation/validation
- `internal/middleware` – Auth, error and logging middlewares
- `internal/schedule` – Appointments handlers, services, repo, utils  
- `pkg/models` – Shared structs and DTOs
- `docker-compose.yaml` – PostgreSQL service setup

---

## ⚙️ Tech Stack

- Go 1.23.4
- PostgreSQL 17
- Docker + Docker Compose
- JWT (`github.com/golang-jwt/jwt/v5`)
- Env vars with `godotenv`

---

## 🚀 Getting Started

```bash
1. Clone the repository
   git clone https://github.com/pablobdss/Backend-Schedule.git
   cd Backend-Schedule

2. Create a `.env` file with your environment variables:
   POSTGRES_USER=your_user
   POSTGRES_PASSWORD=your_password
   POSTGRES_DB=your_db
   POSTGRES_PORT=your_port
   JWT_SECRET=your_secure_jwt_secret

3. Start PostgreSQL with Docker:
   docker-compose up -d

4. Install dependencies:
   go mod tidy

5. Run the project:
   go run cmd/api/main.go
```

---

## 📬 Endpoints (Version 1)

### Public
- `POST /register`  
- `POST /login`  

### Protected (header `Authorization: Bearer <token>`)
- `GET  /dashboard`  
- **Appointments CRUD**  
  - `POST   /schedules`  
  - `GET    /schedules`  
  - `PUT    /schedules/{id}`  
  - `DELETE /schedules/{id}`  

---

## 🐳 Docker

You can run PostgreSQL locally using Docker:
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

## 📍 Roadmap

### 🧱 Version 2 – Infrastructure & Security (coming soon)

- [ ] Dockerize the Go backend
- [ ] JWT expiration validation
- [ ] Basic rate limiting (by IP)
- [ ] Global error handling middleware
- [ ] Structured logs with zerolog or zap
- [ ] Basic tests for main routes
- [ ] Updated README with Docker usage

### ✨ Version 3 – Product Experience (future vision)

- Frontend integration (TypeScript React/Next.js)
- OpenAI-powered chatbot for haircut assistance
- Appointment system with availability validation
- Payment integration (simulated)

---

## Contributing

This project is part of my personal learning journey and portfolio with Go. Feel free to reach out on LinkedIn for feedback or suggestions!
