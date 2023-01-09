package util

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	grpclib "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func zerologUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpclib.UnaryServerInfo, handler grpclib.UnaryHandler) (interface{}, error) {
	l := log.Info().Str("method", info.FullMethod)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		l.Strs("user-agent", md["user-agent"])
	}

	if pr, ok := peer.FromContext(ctx); ok {
		l.Str("ip", pr.Addr.String())
	}

	l.Send()

	return handler(ctx, req)
}

func DefaultServer() *grpclib.Server {
	return grpclib.NewServer(
		grpclib.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			zerologUnaryServerInterceptor,
			grpc_validator.UnaryServerInterceptor(),
			otelgrpc.UnaryServerInterceptor(),
		)),
	)

}
