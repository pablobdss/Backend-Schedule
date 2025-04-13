package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pablobdss/Backend-Schedule/internal/db"
	"github.com/pablobdss/Backend-Schedule/internal/middleware"
	"github.com/pablobdss/Backend-Schedule/internal/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/register", user.RegisterHandler(conn))
	mux.HandleFunc("/login", user.LoginHandler(conn))

	mux.Handle("/dashboard", middleware.AuthMiddleware(http.HandlerFunc(user.DashboardHandler(conn))))

	http.ListenAndServe(":8080", mux)
}
