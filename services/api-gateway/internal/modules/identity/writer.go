package identity

import (
	"context"
)

type Writer interface {
	SetIdentity(ctx context.Context, visitorID string, data Identity) error
	SetVisitorIDRateLimit(ctx context.Context, visitorID string) (int, error)
}
