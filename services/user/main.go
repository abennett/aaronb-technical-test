package main

import (
	"crypto/ed25519"
	"embed"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services"
	"github.com/toggleglobal/aaronb-technical-test/services/user/internal"
)

//go:embed sql/*.sql
var migrations embed.FS

func getKeyPair() (services.KeyPair, error) {
	var kp services.KeyPair
	pubEnc, ok := os.LookupEnv("PUBLIC_KEY")
	if !ok {
		return kp, errors.New("missing PUBLIC_KEY envvar")
	}
	pub, err := base64.StdEncoding.DecodeString(pubEnc)
	if err != nil {
		return kp, err
	}
	privEnc, ok := os.LookupEnv("PRIVATE_KEY")
	if !ok {
		return kp, errors.New("missing PRIVATE_KEY envvar")
	}
	priv, err := base64.StdEncoding.DecodeString(privEnc)
	if err != nil {
		return kp, err
	}
	return services.KeyPair{
		Public:  ed25519.PublicKey(pub),
		Private: ed25519.PrivateKey(priv),
	}, nil
}

func main() {
	l, err := services.NewProductionLogger()
	if err != nil {
		fmt.Println("unable to create logger: " + err.Error())
		os.Exit(1)
	}
	kp, err := getKeyPair()
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
	closer, err := services.SetupTracing("user")
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
	defer closer.Close()
	db, err := services.SetupDB(migrations)
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
	server, err := internal.NewService(l, kp, db)
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
	user := gen.NewUserServiceServer(server, services.BaseHooks(l))
	go func() {
		l.Info("listening on port :"+"8091", zap.String("service", "user-internal"))
		http.ListenAndServe(":8091", services.ServiceWrapper(user))
	}()
	go func() {
		userPub := gen.NewPublicUserServiceServer(server, services.BaseHooks(l))
		l.Info("listening on port :"+"8090", zap.String("service", "user-public"))
		http.ListenAndServe(":8090", services.ServiceWrapper(userPub))
	}()
	l.Info("metrics running", zap.String("port", ":8411"))
	http.ListenAndServe(":8411", promhttp.Handler())
}
