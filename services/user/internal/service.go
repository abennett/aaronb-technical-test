package internal

import (
	"context"
	"fmt"
    "os"
    "database/sql"
    "errors"

    "go.uber.org/zap"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services/user/models"
)

type Service struct {
	q *models.Queries
    l *zap.Logger
    session 
}

func NewService(l *zap.Logger) (*Service, error) {
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
    return &Service{
        q: q,
        l: l,
    }, nil
}

func (s *Service) GetUser(ctx context.Context, req *gen.GetUserReq) (*gen.GetUserResp, error) {
    req.SessionId
	return nil, nil
}

func (s *Service) GetTags(ctx context.Context, req *gen.GetTagsReq) (*gen.GetTagsResp, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid uuid", req.UserId)
	}
	tags, err := s.q.GetUserTags(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tags for user: %w", err)
	}
	resp := &gen.GetTagsResp{
		UserId: req.UserId,
		Tags:   tags,
	}
	return resp, nil
}

func (s *Service) Login(ctx context.Context, req *gen.LoginReq) (*gen.LoginResp, error) {
	return nil, nil
}

func (s *Service) Logout(ctx context.Context, req *gen.LogoutReq) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Service) CreateUser(ctx context.Context, req *gen.CreateUserReq) (*emptypb.Empty, error) {
	return nil, nil
}
