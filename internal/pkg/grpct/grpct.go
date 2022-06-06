package grpct

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"micro-shop/internal/user-srv/svc"
	"runtime/debug"
	"strings"
)

type GrpcInterface interface {
	LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
}

type grpct struct {
	svcCtx *svc.Svc
}

func NewGrpc(svcCtx *svc.Svc) GrpcInterface {
	return &grpct{svcCtx: svcCtx}
}

func (g *grpct) LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if !strings.Contains(info.FullMethod, "Health/Check") {
		if err != nil {
			g.svcCtx.Errorf("[FAILED] GRPC METHOD: %s, ERROR: %s", info.FullMethod, err.Error())
		} else {
			global.Logger.Desugar().WithOptions(zap.WithCaller(false)).Sugar().Infof("[SUCCESS] GRPC METHOD: %s, REQUEST: %v", info.FullMethod, req)
			//global.Logger.Desugar().WithOptions(zap.WithCaller(false)).Sugar().Infof("[SUCCESS] GRPC METHOD: %s, RESPONSE: %v", info.FullMethod, resp)
		}
	}
	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()
	return handler(ctx, req)
}
