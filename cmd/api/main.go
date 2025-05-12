package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/pablobdss/Backend-Schedule/internal/db"
	"github.com/pablobdss/Backend-Schedule/internal/middleware"
	"github.com/pablobdss/Backend-Schedule/internal/schedule"
	"github.com/pablobdss/Backend-Schedule/internal/user"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error to load .env")
	}

	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	r := chi.NewRouter()

	r.Post("/register", user.RegisterHandler(conn))
	r.Post("/login", user.LoginHandler(conn))

	r.With(middleware.AuthMiddleware).Get("/dashboard", user.DashboardHandler(conn))

	r.Route("/schedules", func(r chi.Router) {
		r.With(middleware.AuthMiddleware).Post("/", schedule.CreateScheduleHandler(conn))
		r.With(middleware.AuthMiddleware).Get("/", schedule.GetSchedulesHandler(conn))
		r.With(middleware.AuthMiddleware).Put("/{id}", schedule.UpdateScheduleHandler(conn))
		r.With(middleware.AuthMiddleware).Delete("/{id}", schedule.DeleteScheduleHandler(conn))
	})

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", r)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatalf("failed to load location: %v", err)
	}

	time.Local = loc
}
