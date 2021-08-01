package services

import (
	"context"
	"time"

	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
)

func BaseHooks(l *zap.Logger) *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			method, ok := twirp.MethodName(ctx)
			if ok {
				requestsReceived.WithLabelValues(method).Inc()
			}
			service, _ := twirp.ServiceName(ctx)
			remote := GetRemote(ctx)
			l.Info("request recevied",
				zap.String("service", service),
				zap.String("method", method),
				zap.String("remote_addr", remote),
			)
			return ctx, nil
		},
		ResponseSent: func(ctx context.Context) {
			status, ok := twirp.StatusCode(ctx)
			if ok {
				responsesSent.WithLabelValues(status).Inc()
			}
			start := GetStartTime(ctx)
			method, _ := twirp.MethodName(ctx)
			remote := GetRemote(ctx)
			l.Info("response sent",
				zap.String("method", method),
				zap.String("remote_addr", remote),
				zap.String("status_code", status),
				zap.Duration("request_duration", time.Since(start)),
			)
			return
		},
		Error: func(ctx context.Context, err twirp.Error) context.Context {
			l.Error(err.Error())
			return ctx
		},
	}
}
