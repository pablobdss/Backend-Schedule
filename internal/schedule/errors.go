package schedule

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrInvalidUserID    = errors.New("userID cannot be empty")
	ErrInvalidDate      = errors.New("date cannot be empty")
	ErrInvalidTime      = errors.New("time cannot be empty")
	ErrSlotOccupied     = errors.New("schedule already exists")
	ErrScheduleNotFound = errors.New("schedule not found")
)

func mapScheduleError(w http.ResponseWriter, err error) bool {
	
	if err == nil {
        return false
    }

	switch {
	case errors.Is(err, ErrScheduleNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, ErrInvalidDate), errors.Is(err, ErrInvalidTime):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, ErrInvalidUserID):
		http.Error(w, err.Error(), http.StatusConflict)
	default:
		log.Printf("unexpected schedule error: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
	return err != nil
}
