package identity

func MapToIdentity(identityMap map[string]string) *Identity {
	return &Identity{
		ID:              identityMap["id"],
		IsAuthenticated: identityMap["isAuthenticated"] == "1",
	}
}
