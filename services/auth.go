package services

import (
    "time"
    "net/http"
    "context"
    "crypto/ed25519"

    "github.com/google/uuid"
    "gopkg.in/square/go-jose.v2"
    "gopkg.in/square/go-jose.v2/jwt"
)

const (
    issuer = "aaronb"
    subject = "technical-test"

    authHeader = "Authorization"
)

var idClaimKey idClaim

type Auth struct {
    signer jose.Signer
}

type KeyPair struct {
    Public ed25519.PublicKey
    Private ed25519.PrivateKey
}

type (
    IDClaim struct {
        UserID uuid.UUID
    }

     idClaim struct{}
)

func GetUserIDCtx(ctx context.Context) uuid.UUID {
    id := ctx.Value(idClaimKey).(uuid.UUID)
    return id
}

func MustAuth(pub ed25519.PublicKey, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ah := r.Header.Get(authHeader)
        if ah == "" {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        token, err := jwt.ParseSigned(ah)
        if err != nil {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        claims := new(jwt.Claims)
        idClaim := new(IDClaim)
        if err = token.Claims(pub, claims, idClaim); err != nil {
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        }
        ctx := r.Context()
        ctx = context.WithValue(ctx, idClaimKey, idClaim.UserID)
        r.WithContext(ctx)
        next.ServeHTTP(w, r)
    })
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

func (a *Auth) MintToken(userID uuid.UUID) (string, error) {
    now := time.Now()
    claims := jwt.Claims{
        Subject: subject,
        Issuer: issuer,
        IssuedAt: jwt.NewNumericDate(now),
        Expiry: jwt.NewNumericDate(now.Add(time.Hour)),
    }
    idClaim := &IDClaim{userID}
    raw, err := jwt.Signed(a.signer).Claims(claims).Claims(idClaim).CompactSerialize()
    if err != nil {
        return "", err
    }
    return raw, err
}
