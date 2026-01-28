package service

import (
	"context"

	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1/dto/register"
	"moufube.com/m/internal/modules/authentication/dto/request"
	"moufube.com/m/internal/modules/authentication/dto/response"
)

func (s *Authentication) Register(ctx context.Context, req request.Register) (*response.Register, error) {
	reqPb := &authenticationpb.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	responsePb, err := s.authenticationStub.Register(ctx, reqPb)
	if err != nil {
		return nil, err
	}

	response := &response.Register{
		ID:    responsePb.GetId(),
		Email: responsePb.GetEmail(),
	}

	return response, nil
}
