package server

import (
	"context"
	"log"
	"toolkit/jwtutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// type AccessResource struct {
// 	FullMethod    string
// 	IsRequireAuth bool
// }

type ContextKey string

type RPCInterceptor struct {
	maker         *jwtutil.Maker
	authResources map[string]bool // authResources 包含需要登录授权才能访问的资源
}

func NewRPCInterceptor(maker *jwtutil.Maker, resources map[string]bool) *RPCInterceptor {
	return &RPCInterceptor{
		maker,
		resources,
	}
}

func (r *RPCInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		log.Println(">>>>> Unary Interceptor", info.FullMethod)
		ctx, err = r.requireAuth(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

// wrapper 组合 grpc.ServerStream 的接口和自定义设置context的接口
type wrapper interface {
	grpc.ServerStream
	setContext(context.Context)
}

// wrapp 为了实现 wrapper 接口和重写 Context() 方法
type wrapp struct {
	grpc.ServerStream
	ctx context.Context
}

func newWrapper(stream grpc.ServerStream) wrapper {
	return &wrapp{
		stream,
		stream.Context(),
	}
}

// grpc.ServerStream 重写 Context
func (w *wrapp) Context() context.Context {
	return w.ctx
}

func (w *wrapp) setContext(ctx context.Context) {
	w.ctx = ctx
}

func (r *RPCInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println(">>>>> Stream Interceptor: ", info.FullMethod)

		ctx, err := r.requireAuth(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		// 这里只能获取ctx,不能直接设置
		// stream.Context()

		// 因此裹一层
		w := newWrapper(stream)
		w.setContext(ctx)

		return handler(srv, w)
	}
}

func (r *RPCInterceptor) requireAuth(ctx context.Context, fullMethod string) (context.Context, error) {
	_, ok := r.authResources[fullMethod]
	if !ok {
		// 没有权限要求
		return ctx, nil
	}
	// 需要登录用户才能访问
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "未授权")
	}

	tokens := md["authorization"]
	if len(tokens) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "token 不存在")
	}

	token := tokens[0]
	payload, err := r.maker.ParseToken(token)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "token 无效： %v", err)
	}

	ctx = context.WithValue(ctx, ContextKey("user_id"), payload.UID)

	return ctx, nil

}
