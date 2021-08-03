package services

import (
	"context"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
)

func BaseHooks(l *zap.Logger) *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			span, ctx := opentracing.StartSpanFromContext(ctx, "request_routed")
			defer span.Finish()
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
			span, _ := opentracing.StartSpanFromContext(ctx, "response_sent")
			defer span.Finish()
			start := GetStartTime(ctx)
			method, _ := twirp.MethodName(ctx)
			remote := GetRemote(ctx)
			dur := time.Since(start)
			l.Info("response sent",
				zap.String("method", method),
				zap.String("remote_addr", remote),
				zap.String("status_code", status),
				zap.Duration("request_duration", dur),
			)
			responseLatency.WithLabelValues(status, method).Observe(dur.Seconds())
			return
		},
		Error: func(ctx context.Context, err twirp.Error) context.Context {
			l.Error(err.Error())
			return ctx
		},
	}
}
