package service

import (
	"context"

	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1/dto/login"
	"moufube.com/m/internal/modules/authentication/dto/request"
	"moufube.com/m/internal/modules/authentication/dto/response"
)

func (s *Authentication) Login(ctx context.Context, req request.Login) (*response.Login, error) {
	reqPb := &authenticationpb.LoginRequest{
		Identifier: req.Identifier,
		Password:   req.Password,
	}

	responsePb, err := s.authenticationStub.Login(ctx, reqPb)
	if err != nil {
		return nil, err
	}

	loginResponse := &response.Login{
		Message: responsePb.GetMessage(),
		User: response.User{
			ID:         responsePb.GetUser().GetId(),
			IsVerified: responsePb.GetUser().GetIsVerified(),
			Email:      responsePb.GetUser().GetEmail(),
		},
		Tokens: response.Tokens{
			AccessToken:  responsePb.GetTokens().GetAccessToken(),
			RefreshToken: responsePb.GetTokens().GetRefreshToken(),
		},
	}

	return loginResponse, nil
}
