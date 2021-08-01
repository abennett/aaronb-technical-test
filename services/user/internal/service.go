package internal

import (
	"context"
	"fmt"
    "os"
    "database/sql"
    "errors"

    "go.uber.org/zap"
	"github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services"
	"github.com/toggleglobal/aaronb-technical-test/services/user/models"
)

type Service struct {
	q *models.Queries
    kp services.KeyPair
    auth *services.Auth
    l *zap.Logger
}

func NewService(l *zap.Logger, kp services.KeyPair) (*Service, error) {
	pgConn, ok := os.LookupEnv("PG_CONN")
	if !ok {
		return nil, errors.New("PG_CONN not available")
	}
	db, err := sql.Open("postgres", pgConn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	q := models.New(db)
    auth, err := services.NewAuth(kp.Public, kp.Private)
    if err != nil {
        return nil, err
    }
    return &Service{
        q: q,
        auth: auth,
        kp: kp,
        l: l,
    }, nil
}

func (s *Service) GetUserTags(ctx context.Context, req *gen.GetUserTagsReq) (*gen.GetUserTagsResp, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid uuid", req.UserId)
	}
	tags, err := s.q.GetUserTags(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tags for user: %w", err)
	}
	resp := &gen.GetUserTagsResp{
		UserId: req.UserId,
		Tags:   tags,
	}
	return resp, nil
}

func (s *Service) CreateUser(ctx context.Context, req *gen.CreateUserReq) (*emptypb.Empty, error) {
    empty := &emptypb.Empty{}
    if req.Username == "" {
        return empty, errors.New("username cannot be blank")
    }
    if req.Password == "" {
        return empty, errors.New("password cannot be blank")
    }
    hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return empty, errors.New("failed to hash password")
    }
    params := models.CreateUserParams{
        Name: req.Username,
        Password: hash,
    }
    if _, err = s.q.CreateUser(ctx, params); err != nil {
        return empty, errors.New("unable to create new user")
    }
	return empty, nil
}

func (s *Service) GetPublicKey(ctx context.Context, _ *emptypb.Empty) (*gen.GetPublicKeyResp, error) {
    resp := &gen.GetPublicKeyResp{
        PublicKey: []byte(s.kp.Public),
    }
    return resp, nil
}

func (s *Service) Login(ctx context.Context, req *gen.LoginReq) (*gen.LoginResp, error) {
    if req.Username == "" || req.Password == "" {
        return nil, errors.New("username and password must not be empty")
    }
    user, err := s.q.GetUserByName(ctx, req.Username)
    if err != nil {
        return nil, err
    }
    if err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
        return nil, err
    }
    token, err := s.auth.MintToken(user.ID)
    if err != nil {
        return nil, err
    }
    resp := &gen.LoginResp{
        Token: token,
    }
    return resp, nil
}
