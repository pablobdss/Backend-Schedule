package schedule

type ScheduleRequest struct {
	Date          string `json:"date"`
	ScheduledTime string `json:"time"`
}

type ScheduleResponse struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	Date          string `json:"date"`
	ScheduledTime string `json:"time"`
}
