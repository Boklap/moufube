package identity

import (
	"context"
)

type Reader interface {
	GetIdentityByID(ctx context.Context, visitorID string) (map[string]string, error)
}
