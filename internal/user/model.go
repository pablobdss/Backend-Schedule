package user

type User struct {
	ID       string `json:"id"` 
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DashboardResponse struct {
	UserID string `json:"user_id"`
	Status string `json:"status"`
}