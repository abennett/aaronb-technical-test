package main

import (
	"context"
	"crypto/ed25519"
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services"
	"github.com/toggleglobal/aaronb-technical-test/services/news/internal"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

func main() {
	l, err := services.NewProductionLogger()
	if err != nil {
		fmt.Println("unable to create logger: " + err.Error())
		os.Exit(1)
	}
	closer, err := services.SetupTracing("news")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer closer.Close()
	db, err := services.SetupDB(sqlFiles)
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
	// setup user service client
	userAddr, ok := os.LookupEnv("USER_SRV")
	if !ok {
		l.Error("USER_SRV not available")
		os.Exit(1)
	}
	users := gen.NewUserServiceProtobufClient(userAddr, &http.Client{})
	server, err := internal.NewService(l, db, users)
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
	newsSrv := gen.NewNewsServiceServer(server, services.BaseHooks(l))
	newsHandler := services.ServiceWrapper(newsSrv)
	pub, err := users.GetPublicKey(context.Background(), nil)
	if err != nil {
		l.Error("failed to fetch pubic key from user service", zap.Error(err))
		os.Exit(1)
	}
	pubKey := ed25519.PublicKey(pub.PublicKey)
	newsHandler = services.MustAuth(pubKey, newsHandler)
	go func() {
		l.Info("listening on port :"+"8100", zap.String("service", "news"))
		if err := http.ListenAndServe(":8100", newsHandler); err != nil {
			l.Error("news service returned an error", zap.Error(err))
		}
	}()
	l.Info("metrics running", zap.String("port", ":8411"))
	http.ListenAndServe(":8411", promhttp.Handler())
}
