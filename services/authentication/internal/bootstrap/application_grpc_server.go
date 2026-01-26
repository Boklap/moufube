package bootstrap

import "moufube.com/m/internal/application/grpc_server/user"

type ApplicationGRPCServer struct {
	UserGRPCServer *user.GRPCServerImpl
}

func InitApplicationGRPCServer(writer *Writer, reader *Reader) *ApplicationGRPCServer {
	userGRPCServer := user.NewGRPCServerImpl(reader.UserReader, writer.UserWriter)

	return &ApplicationGRPCServer{
		UserGRPCServer: userGRPCServer,
	}
}
