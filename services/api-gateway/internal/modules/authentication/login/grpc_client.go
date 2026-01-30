package login

import (
	"context"

	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1"
	loginpb "moufube.com/m/internal/generated/pb/authentication/v1/login"
)

type GRPCClient struct {
	authenticationClient authenticationpb.AuthenticationClient
}

func NewGRPCClient(authenticationClient authenticationpb.AuthenticationClient) *GRPCClient {
	return &GRPCClient{
		authenticationClient: authenticationClient,
	}
}

func (c *GRPCClient) Login(ctx context.Context, req *loginpb.Request) (*loginpb.Response, error) {
	return c.authenticationClient.Login(ctx, req)
}
