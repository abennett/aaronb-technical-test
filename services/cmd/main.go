package main

import (
    "fmt"
    "log"
    "encoding/base64"
    "crypto/rand"
    "crypto/ed25519"
)

func main() {
    pub, priv, err := ed25519.GenerateKey(rand.Reader)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(base64.StdEncoding.EncodeToString([]byte(priv)))
    fmt.Println(base64.StdEncoding.EncodeToString([]byte(pub)))
}
