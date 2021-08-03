package services

import (
	"context"
	"crypto/ed25519"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

const (
	issuer  = "aaronb"
	subject = "technical-test"

	authHeader = "Authorization"
)

var idClaimKey idClaim

type Auth struct {
	signer jose.Signer
}

type KeyPair struct {
	Public  ed25519.PublicKey
	Private ed25519.PrivateKey
}

type IDClaim struct {
	UserID uuid.UUID
}

type idClaim struct{}

func GetUserIDCtx(ctx context.Context) uuid.UUID {
	if id, ok := ctx.Value(idClaimKey).(uuid.UUID); ok {
		return id
	}
	return uuid.Nil
}

func InsertAuthHeader(ctx context.Context, token string) (context.Context, error) {
	header := make(http.Header)
	header.Set(authHeader, token)
	return twirp.WithHTTPRequestHeaders(ctx, header)
}

func InjectAuth(l *zap.Logger, pub ed25519.PublicKey, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ah := r.Header.Get(authHeader)
		if ah == "" {
			l.Error("no Authorization header provided")
			next.ServeHTTP(w, r)
			return
		}
		token, err := jwt.ParseSigned(ah)
		if err != nil {
			l.Error("failed to parse token", zap.Error(err))
			next.ServeHTTP(w, r)
			return
		}
		claims := new(jwt.Claims)
		idClaim := new(IDClaim)
		if err = token.Claims(pub, claims, idClaim); err != nil {
			l.Error("failed to validate claims", zap.Error(err))
			next.ServeHTTP(w, r)
			return
		}
		ctx := r.Context()
		l.Info("user authenticated", zap.String("id", idClaim.UserID.String()))
		ctx = context.WithValue(ctx, idClaimKey, idClaim.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func MustAuth(pub ed25519.PublicKey, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ah := r.Header.Get(authHeader)
		zap.L().Info("auth header", zap.String("header", ah))
		if ah == "" {
			twirp.WriteError(w, twirp.NewError(twirp.Unauthenticated, "authentication required"))
			return
		}
		token, err := jwt.ParseSigned(ah)
		if err != nil {
			twirp.WriteError(w, twirp.NewError(twirp.Unauthenticated, "failed to parse token"))
			return
		}
		claims := new(jwt.Claims)
		idClaim := new(IDClaim)
		if err = token.Claims(pub, claims, idClaim); err != nil {
			twirp.WriteError(w, twirp.NewError(twirp.Unauthenticated, "fail to validate claims"))
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, idClaimKey, idClaim.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func NewAuth(pub, priv []byte) (*Auth, error) {
	key := jose.SigningKey{
		Algorithm: jose.EdDSA,
		Key:       ed25519.PrivateKey(priv),
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
		Subject:  subject,
		Issuer:   issuer,
		IssuedAt: jwt.NewNumericDate(now),
		Expiry:   jwt.NewNumericDate(now.Add(time.Hour)),
	}
	idClaim := &IDClaim{userID}
	raw, err := jwt.Signed(a.signer).Claims(claims).Claims(idClaim).CompactSerialize()
	if err != nil {
		return "", err
	}
	return raw, err
}
