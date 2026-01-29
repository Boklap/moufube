package register

import (
	"context"

	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1/register"
)

type Service struct {
	grpcClient *GRPCClient
}

func NewService(grpcClient *GRPCClient) *Service {
	return &Service{
		grpcClient: grpcClient,
	}
}

func (s *Service) Register(ctx context.Context, req *Request) (*Response, error) {
	reqPb := &authenticationpb.Request{
		Email:    req.Email,
		Password: req.Password,
	}

	responsePb, err := s.grpcClient.authenticationClient.Register(ctx, reqPb)
	if err != nil {
		return nil, err
	}

	response := &Response{
		ID:    responsePb.GetId(),
		Email: responsePb.GetEmail(),
	}

	return response, nil
}
