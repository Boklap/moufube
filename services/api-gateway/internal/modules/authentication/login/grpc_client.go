package login

import authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1"

type GRPCClient struct {
	authenticationClient authenticationpb.AuthenticationClient
}

func NewGRPCClient(authenticationClient authenticationpb.AuthenticationClient) *GRPCClient {
	return &GRPCClient{
		authenticationClient: authenticationClient,
	}
}
