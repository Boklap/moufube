package service

import authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1/contract"

func NewAuthenticationService(authStub authenticationpb.AuthenticationClient) *Authentication {
	return &Authentication{
		authenticationStub: authStub,
	}
}
