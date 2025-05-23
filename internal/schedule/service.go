package schedule

import (
	"database/sql"
	"time"
)

func InsertScheduleService(db *sql.DB, userID, date, scheduledTime string) (*ScheduleResponse, error) {
	parsedDate, parsedTime, err := ParseDateTime(date, scheduledTime)
	if err != nil {
		return nil, err
	}

	err = ValidateScheduleInput(userID, parsedDate, parsedTime)
	if err != nil {
		return nil, err
	}

	if IsInThePast(parsedDate, parsedTime) {
		return nil, ErrPastDateTime
	}

	combined := time.Date(
		parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
		parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0,
		time.Local,
	)

	if !IsWithinBusinessHours(combined) {
		return nil, ErrOutsideBusinessHours
	}

	err = EnsureSlotIsFree(db, parsedDate, parsedTime)
	if err != nil {
		return nil, err
	}

	id, err := InsertSchedule(db, userID, parsedDate, parsedTime)
	if err != nil {
		return nil, err
	}

	return &ScheduleResponse{
		ID:            id,
		UserID:        userID,
		Date:          date,
		ScheduledTime: scheduledTime,
	}, nil
}

func UpdateScheduleService(db *sql.DB, id, userID, date, scheduledTime string) (*ScheduleResponse, error) {
	parsedDate, parsedTime, err := ParseDateTime(date, scheduledTime)
	if err != nil {
		return nil, err
	}

	if err = ValidateScheduleInput(userID, parsedDate, parsedTime); err != nil {
		return nil, err
	}

	if IsInThePast(parsedDate, parsedTime) {
		return nil, ErrPastDateTime
	}

	if err = UpdateSchedule(db, id, parsedDate, parsedTime); err != nil {
		return nil, err
	}

	schedule, err := FindScheduleByID(db, id, userID)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func DeleteScheduleService(db *sql.DB, id, userID string) error {
	if _, err := FindScheduleByID(db, id, userID); err != nil {
		return err
	}

	if err := DeleteSchedule(db, id, userID); err != nil {
		return err
	}

	return nil
}
