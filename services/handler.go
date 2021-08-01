package services

import (
	"context"
	"net/http"
	"time"
)

type (
	remoteKey    struct{}
	startTimeKey struct{}

	Middleware func(http.Handler) http.Handler
)

var (
	rk  = remoteKey{}
	stk = startTimeKey{}

	middlewareChain = []Middleware{RemoteHostHandler, StartTimeHandler}
)

func GetRemote(ctx context.Context) string {
	remote := ctx.Value(rk).(string)
	return remote
}

func GetStartTime(ctx context.Context) time.Time {
	start := ctx.Value(stk).(time.Time)
	return start
}

func ServiceWrapper(handler http.Handler) http.Handler {
	for _, h := range middlewareChain {
		handler = h(handler)
	}
	return handler
}

func StartTimeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := context.WithValue(r.Context(), stk, start)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func RemoteHostHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xForwarded := r.Header.Get("X-Forwarded-For")
		if xForwarded != "" {
			r.RemoteAddr = xForwarded
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, rk, r.RemoteAddr)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
