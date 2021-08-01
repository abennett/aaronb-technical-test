package services

import (
    "time"
    //"crypto/rand"
    "crypto/ed25519"

    "gopkg.in/square/go-jose.v2"
    "gopkg.in/square/go-jose.v2/jwt"
)

const (
    issuer = "aaronb"
    subject = "technical-test"
)

var (
    privateKey *ed25519.PrivateKey
    publicKey *ed25519.PublicKey
)

type Auth struct {
    signer jose.Signer
}

func NewAuth(pub, priv []byte) (*Auth, error) {
    key := jose.SigningKey{
        Algorithm: jose.EdDSA,
        Key: ed25519.PrivateKey(priv),
    }
    opts := &jose.SignerOptions{}
    sig, err := jose.NewSigner(key, opts.WithType("JWT"))
    if err != nil {
        return nil, err
    }
    return &Auth{
        signer: sig,
    }, nil
}

func (a *Auth) MintToken() (string, error) {
    now := time.Now()
    claims := jwt.Claims{
        Subject: subject,
        Issuer: issuer,
        IssuedAt: jwt.NewNumericDate(now),
        Expiry: jwt.NewNumericDate(now.Add(time.Hour)),
    }
    raw, err := jwt.Signed(a.signer).Claims(claims).CompactSerialize()
    if err != nil {
        return "", err
    }
    return raw, err
}
