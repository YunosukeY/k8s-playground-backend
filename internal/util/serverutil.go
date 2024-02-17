package util

import (
	"context"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
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
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
			zerologUnaryServerInterceptor,
			grpc_validator.UnaryServerInterceptor(),
		),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

}
