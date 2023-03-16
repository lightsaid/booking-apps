package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	accessToken     string
	authFullMethods map[string]bool // 需要拦截的添加token的 fullmethod
}

func NewAuthInterceptor(token string, methods map[string]bool) *AuthInterceptor {
	return &AuthInterceptor{
		accessToken:     token,
		authFullMethods: methods,
	}
}
func (a *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context, method string,
		req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		log.Println(">>>>> Unary Interceptor: ", method)

		// 需要token
		if a.authFullMethods[method] {
			newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", a.accessToken)
			log.Println("-----> newCtx")
			return invoker(newCtx, method, req, reply, cc, opts...)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (a *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc,
		cc *grpc.ClientConn, method string,
		streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

		log.Println(">>>>> Stream Interceptor: ", method)

		if a.authFullMethods[method] {
			newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", a.accessToken)
			return streamer(newCtx, desc, cc, method, opts...)
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}
