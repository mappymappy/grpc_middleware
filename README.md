# gRPC Middleware For Golang
[![GoDoc](http://godoc.org/github.com/mappymappy/grpc_middleware?status.svg)](http://godoc.org/github.com/mappymappy/grpc_middleware)

common middlewares(such as authnetication,panic_recover,etc) written by golang.

## install

go get `github.com/mappymappy/grpc_middleware`

## Usage

```
	yourAuthenticator := newYourAuthentiator() // use whatever you like
	panicHandler := newyourPanicHandler() // use whatever you like
	handlers := []panci_recover.HandlerInterface{panicHandler}
	// When using streamingRPC
	streamOpts := grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			authentication.NewMiddleware(yourAuthentiator).StreamServerInterceptor,
			panic_recover.NewMiddleware(hadnlers).StreamServerInterceptor,
	))
	// when using UnaryRPC
	unaryOpts := grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			authentication.NewMiddleware(yourAuthentiator).UnaryServerInterceptor,
			panic_recover.NewMiddleware(handlers).UnaryServerInterceptor,
	))
	opts := []grpc.ServerOption{
		streamMiddlewares,
		unaryMiddlewares,
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time: 100 * time.Second,
		}),
	}
	grpcServer := grpc.NewServer(opts)
	grpcServer.Serve("localhost:8080")
```

## Author
[marnie_ms4](https://github.com/mappymappy?tab=repositories)
