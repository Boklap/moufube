package outbound

import (
	"context"

	"moufube.com/m/internal/modules/authentication/dto/request"
	"moufube.com/m/internal/modules/authentication/dto/response"
)

type AuthenticationService interface {
	Register(ctx context.Context, req request.Register) (*response.Register, error)
}
