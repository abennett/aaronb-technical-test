package main

import (
	"fmt"
	"os"
    "errors"
	"net/http"
    "encoding/base64"
    "crypto/ed25519"

	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services"
	"github.com/toggleglobal/aaronb-technical-test/services/user/internal"
)

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
        Public: ed25519.PublicKey(pub),
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
	server, err := internal.NewService(l, kp)
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
	user := gen.NewUserServiceServer(server, services.BaseHooks(l))
	l.Info("listening on port :" + "8090")
	http.ListenAndServe(":8090", services.ServiceWrapper(user))
}
