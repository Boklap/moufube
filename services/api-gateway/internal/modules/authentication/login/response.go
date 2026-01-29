package login

type Response struct {
	Message string `json:"message"`
	User    User   `json:"user"`
	Tokens  Tokens `json:"tokens"`
}

type User struct {
	ID         string `json:"id"`
	IsVerified bool   `json:"is_verified"`
	Email      string `json:"email"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
