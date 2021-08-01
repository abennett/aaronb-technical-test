package main

import (
	"os"
	"fmt"
    "context"
	"net/http"
    "crypto/ed25519"

	"go.uber.org/zap"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services"
	"github.com/toggleglobal/aaronb-technical-test/services/news/internal"
)

func main() {
	c := zap.NewProductionConfig()
	c.DisableCaller = true
	c.DisableStacktrace = true
	l, err := c.Build()
	if err != nil {
		fmt.Println("unable to create logger: " + err.Error())
		os.Exit(1)
	}
	server, err := internal.NewService(l)
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
	newsSrv := gen.NewNewsServiceServer(server, services.BaseHooks(l))
	newsHandler := services.ServiceWrapper(newsSrv)
    pub, err := server.Users.GetPublicKey(context.Background(), nil)
    if err != nil {
        l.Error("failed to fetch pubic key from user service", zap.Error(err))
        os.Exit(1)
    }
    pubKey := ed25519.PublicKey(pub.PublicKey)
    newsHandler = services.MustAuth(pubKey, newsHandler)
	mux := http.NewServeMux()
	mux.Handle("/", newsHandler)
	mux.Handle("/metrics", promhttp.Handler())
	l.Info("listening on port :" + "8100")
	http.ListenAndServe(":8100", mux)
}
