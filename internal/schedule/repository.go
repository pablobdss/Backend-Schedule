package schedule

import (
	"database/sql"
	"time"
)

func EnsureSlotIsFree(db *sql.DB, date, scheduledTime time.Time) error {

	var id string

	err := db.QueryRow("SELECT id FROM schedules WHERE date = $1 AND time = $2", date, scheduledTime).Scan(&id)
	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		return err
	}

	return ErrSlotOccupied
}

func InsertSchedule(db *sql.DB, userID string, date, scheduledTime time.Time) (string, error) {
	var id string

	if err := db.QueryRow(`
		INSERT INTO schedules 
		(user_id, date, time)
		VALUES ($1, $2, $3)
		RETURNING id`, 
	userID, date, scheduledTime,
	).Scan(&id); 
	err != nil {
		return "", err
	}

	return id, nil
}

func GetScheduleByUserID(db *sql.DB, userID string) ([]ScheduleResponse, error) {
	var schedules []ScheduleResponse

	rows, err := db.Query(`
		SELECT 
			id, 
			user_id, 
			TO_CHAR(date, 'YYYY-MM-DD') as date, 
			TO_CHAR(time, 'HH24:MI:SS') as time 
		FROM schedules WHERE user_id = $1 ORDER BY date ASC, time ASC`,
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, user_id, date, time string
		if err := rows.Scan(
			&id,
			&user_id,
			&date,
			&time); err != nil {
			return nil, err
		}

		schedules = append(schedules, ScheduleResponse{
			ID:            id,
			UserID:        user_id,
			Date:          date,
			ScheduledTime: time,
		})
	}

	return schedules, nil
}

func FindScheduleByID(db *sql.DB, id, userID string) (*ScheduleResponse, error) {

	row := db.QueryRow(`
		SELECT 
			id, 
			user_id, 
			TO_CHAR(date,'YYYY-MM-DD'), 
			TO_CHAR(time,'HH24:MI:SS') 
		FROM schedules WHERE id = $1 AND user_id = $2`,
		id, userID)

	var resp ScheduleResponse

	if err := row.Scan(
		&resp.ID,
		&resp.UserID,
		&resp.Date,
		&resp.ScheduledTime); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrScheduleNotFound
		}
		return nil, err
	}
	return &resp, nil
}

func UpdateSchedule(db *sql.DB, id string, date, scheduledTime time.Time) error {
	result, err := db.Exec(`
	UPDATE schedules 
		SET date = $1,
			time = $2 
	WHERE id = $3`,
		date, scheduledTime, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrScheduleNotFound
	}

	return nil
}

func DeleteSchedule(db *sql.DB, id, userID string) error {
	result, err := db.Exec(`
	DELETE 
	FROM schedules 
	WHERE id = $1
	AND user_id = $2`,
		id, userID)
	if err != nil {
		return err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsDeleted == 0 {
		return ErrScheduleNotFound
	}

	return nil
}
