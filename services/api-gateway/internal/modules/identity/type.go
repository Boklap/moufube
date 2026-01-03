package identity

type Identity struct {
	ID              string `redis:"id"`
	IsAuthenticated bool   `redis:"isAuthenticated"`
}
