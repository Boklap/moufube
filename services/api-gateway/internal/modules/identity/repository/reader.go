package repository

import (
	"context"

	"moufube.com/m/internal/modules/identity"
)

type IdentityReader interface {
	GetIdentityByID(ctx context.Context, visitorID string) (identity.Identity, error)
}
