package schedule

import "time"

func ValidateScheduleInput(userID string, date, scheduledTime time.Time) error {
	if userID == "" {
		return ErrInvalidUserID
	}

	if date.IsZero() {
		return ErrInvalidDate
	}

	if scheduledTime.IsZero() {
		return ErrInvalidTime
	}

	return nil
}

func ParseDateTime(date, scheduledTime string) (time.Time, time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	parsedTime, err := time.Parse("15:04:05", scheduledTime)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return parsedDate, parsedTime, nil
}
