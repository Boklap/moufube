package login

type Response struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

type User struct {
	ID         string `json:"id"`
	IsVerified bool   `json:"is_verified"`
	Email      string `json:"email"`
}
