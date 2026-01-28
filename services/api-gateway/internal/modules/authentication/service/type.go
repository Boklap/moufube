package service

import authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1/contract"

type Authentication struct {
	authenticationStub authenticationpb.AuthenticationClient
}
