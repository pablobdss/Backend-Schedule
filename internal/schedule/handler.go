package schedule

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pablobdss/Backend-Schedule/internal/middleware"
)

func CreateScheduleHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID, ok := middleware.GetUserID(r.Context())
		if !ok {
			http.Error(w, "userID not found in context", http.StatusUnauthorized)
			return
		}

		var req ScheduleRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("Erro %v", err)
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}

		newSchedule, err := InsertScheduleService(db, userID, req.Date, req.ScheduledTime)
		if mapScheduleError(w, err) {
			return
		}

		resp := ScheduleResponse{
			ID:            newSchedule.ID,
			UserID:        newSchedule.UserID,
			Date:          newSchedule.Date,
			ScheduledTime: newSchedule.ScheduledTime,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	}

}

func GetSchedulesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID, ok := middleware.GetUserID(r.Context())
		if !ok {
			http.Error(w, "UserID not found in context", http.StatusUnauthorized)
			return
		}

		schedules, err := GetScheduleByUserID(db, userID)
		if err != nil {
			log.Printf("Erro %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) //

		if err := json.NewEncoder(w).Encode(schedules); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}

	}
}

func UpdateScheduleHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, ok := middleware.GetUserID(r.Context())
		if !ok {
			http.Error(w, "UserID not found in context", http.StatusUnauthorized)
			return
		}

		var req ScheduleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		updated, err := UpdateScheduleService(db, id, userID, req.Date, req.ScheduledTime)
		if mapScheduleError(w, err) {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)
	}
}

func DeleteScheduleHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, ok := middleware.GetUserID(r.Context())
		if !ok {
			http.Error(w, "UserID not found in context", http.StatusUnauthorized)
			return
		}

		err := DeleteScheduleService(db, id, userID)
		if mapScheduleError(w, err) {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
