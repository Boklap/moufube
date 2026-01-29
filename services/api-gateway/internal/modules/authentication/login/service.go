package login

import (
	"context"

	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1/login"
)

type Service struct {
	grpcClient *GRPCClient
}

func NewService(grpcClient *GRPCClient) *Service {
	return &Service{
		grpcClient: grpcClient,
	}
}

func (s *Service) Login(ctx context.Context, req *Request) (*Response, error) {
	reqPb := &authenticationpb.Request{
		Identifier: req.Identifier,
		Password:   req.Password,
	}

	resPb, err := s.grpcClient.authenticationClient.Login(ctx, reqPb)
	if err != nil {
		return nil, err
	}

	res := &Response{
		Message: resPb.GetMessage(),
		User: User{
			ID:         resPb.GetUser().GetId(),
			IsVerified: resPb.GetUser().GetIsVerified(),
			Email:      resPb.GetUser().GetEmail(),
		},
		Tokens: Tokens{
			AccessToken:  resPb.GetTokens().GetAccessToken(),
			RefreshToken: resPb.GetTokens().GetRefreshToken(),
		},
	}

	return res, nil
}
