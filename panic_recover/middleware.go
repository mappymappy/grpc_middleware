package panic_recover

import (
	"context"

	"google.golang.org/grpc"
)

type Middleware struct {
	handlers []HandlerInterface
}

type HandlerInterface interface {
	Handle(context.Context, interface{})
}

func NewMiddleware(handlers []HandlerInterface) *Middleware {
	return &Middleware{
		handlers,
	}
}

func (m *Middleware) StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func(stream grpc.ServerStream) {
		if r := recover(); r != nil {
			for _, handler := range handlers {
				handler.Handle(ctx, r)
			}
		}
	}(ss)
	return handler(srv, ss)
}

func (m *Middleware) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			for _, handler := range handlers {
				handler.Handle(ctx, r)
			}
		}
	}()
	return handler(ctx, req)
}
