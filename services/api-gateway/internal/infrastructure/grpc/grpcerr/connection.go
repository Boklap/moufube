package grpcerr

import "errors"

var ErrGRPCConnectionFailed = errors.New("fail to connect to GRPC server")
