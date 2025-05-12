package schedule

import (
	"time"
)

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
	parsedDate, err := time.ParseInLocation("2006-01-02", date, time.Local)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	parsedTime, err := time.ParseInLocation("15:04:05", scheduledTime, time.Local)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return parsedDate, parsedTime, nil
}

func IsInThePast(parsedDate, parsedTime time.Time) bool {
	now := time.Now()
	combined := time.Date(
		parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
		parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0,
		time.Local,
	)

	return combined.Before(now)
}

func IsWithinBusinessHours(parsedTime time.Time) bool {
    h := parsedTime.Hour()
    return h >= 8 && h < 18
}
