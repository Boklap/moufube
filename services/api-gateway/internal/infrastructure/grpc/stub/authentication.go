package stub

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"moufube.com/m/internal/apperr"
	"moufube.com/m/internal/config"
	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1"
)

func InitAuthenticationStub(
	cfg *config.Config,
) (authenticationpb.AuthenticationClient, *grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	serverAddress := fmt.Sprintf("%s:%d", cfg.GRPCAuthenticationHost, cfg.GRPCAuthenticationPort)

	conn, err := grpc.NewClient(serverAddress, opts...)
	if err != nil {
		grpcErr := apperr.NewGRPCError(apperr.ErrGRPCConnectionFailed, "authentication")
		return nil, nil, grpcErr
	}

	client := authenticationpb.NewAuthenticationClient(conn)

	return client, conn, nil
}
