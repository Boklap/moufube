package repository

import (
	"context"

	"moufube.com/m/internal/modules/identity"
)

type IdentityWriter interface {
	SetIdentity(ctx context.Context, visitorID string, data identity.Identity) error
}
