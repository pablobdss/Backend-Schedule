package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pablobdss/Backend-Schedule/internal/auth"
	"github.com/pablobdss/Backend-Schedule/internal/middleware"
	"github.com/pablobdss/Backend-Schedule/pkg/models"
)

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req models.RegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}

		newUser, err := RegisterUser(req.Name, req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := CreateUser(db, newUser); err != nil {

			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		resp := models.RegisterResponse{
			ID:    newUser.ID,
			Name:  newUser.Name,
			Email: newUser.Email,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}

	}

}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req models.LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}

		login, err := LoginUser(db, req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := auth.GenerateToken(login.ID, login.Email)
		if err != nil {
			http.Error(w, "Error to generate token", http.StatusInternalServerError)
			return
		}

		login.Token = token

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(login); err != nil {
			http.Error(w, "Error Coding response", http.StatusInternalServerError)
			return
		}
	}
}

func DashboardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := middleware.GetUserID(r.Context())
		if !ok {
			http.Error(w, "userID not found in context", http.StatusUnauthorized)
			return
		}
		log.Println("Authenticated user ID:", userID)

		resp := DashboardResponse{
			UserID: userID,
			Status: "ok",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
