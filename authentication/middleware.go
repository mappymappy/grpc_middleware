package authentication

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Middleware struct {
	authenticator AuthenticatorInterface
}

type AuthenticatorInterface interface {
	AuthenticateByMD(context.Context, metadata.MD) (context.Context, error)
}

func NewMiddleware(authenticator AuthenticatorInterface) *Middleware {
	return &Middleware{
		authenticator: authenticator,
	}
}

// implement StreamServerInterceptor. more information -> @see google.golang.org/grpc/intercepter.go
func (m *Middleware) StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := ss.Context()
	grpcMeta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.InvalidArgument, "grpcMeta cast error")
	}
	newCtx, err := m.authenticator.AuthenticateByMD(ctx, grpcMeta)
	if err != nil {
		return err
	}
	wrap := grpc_middleware.WrapServerStream(ss)
	wrap.WrappedContext = newCtx

	return handler(srv, wrap)
}

func (m *Middleware) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	grpcMeta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "grpcMeta cast error")
	}
	newCtx, err := m.authenticator.AuthenticateByMD(ctx, grpcMeta)
	if err != nil {
		return nil, err
	}
	return handler(newCtx, req)
}
