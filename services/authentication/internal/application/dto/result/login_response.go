package result

type LoginUser struct {
	Message      string
	ID           string
	IsVerified   bool
	Email        string
	AccessToken  string
	RefreshToken string
}
