package internal

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services"
	"github.com/toggleglobal/aaronb-technical-test/services/user/models"
)

type Service struct {
	q      *models.Queries
	kp     services.KeyPair
	auth   *services.Auth
	l      *zap.Logger
	tracer opentracing.Tracer
}

func NewService(l *zap.Logger, kp services.KeyPair, db *sql.DB) (*Service, error) {
	q := models.New(db)
	auth, err := services.NewAuth(kp.Public, kp.Private)
	if err != nil {
		return nil, err
	}
	return &Service{
		q:      q,
		auth:   auth,
		kp:     kp,
		l:      l,
		tracer: opentracing.GlobalTracer(),
	}, nil
}

func (s *Service) GetUserTags(ctx context.Context, req *gen.GetUserTagsReq) (*gen.GetUserTagsResp, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, twirp.NewError(twirp.InvalidArgument, "provided ID is not a valid uuid")
	}
	tags, err := s.q.GetUserTags(ctx, id)
	if err != nil {
		return nil, twirp.NewError(twirp.Internal, "failed to fetch tags for uuid: "+req.UserId)
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
		return empty, twirp.NewError(twirp.InvalidArgument, "username cannot be blank")
	}
	if req.Password == "" {
		return empty, twirp.NewError(twirp.InvalidArgument, "password cannot be blank")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return empty, twirp.NewError(twirp.Internal, "failed to hash password")
	}
	params := models.CreateUserParams{
		Name:     req.Username,
		Password: hash,
	}
	if _, err = s.q.CreateUser(ctx, params); err != nil {
		return empty, twirp.NewError(twirp.Internal, "unable to create new user")
	}
	return empty, nil
}

func (s *Service) GetPublicKey(ctx context.Context, _ *emptypb.Empty) (*gen.GetPublicKeyResp, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetPublicKey")
	defer span.Finish()
	resp := &gen.GetPublicKeyResp{
		PublicKey: []byte(s.kp.Public),
	}
	return resp, nil
}

func (s *Service) Login(ctx context.Context, req *gen.LoginReq) (*gen.LoginResp, error) {
	if req.Username == "" || req.Password == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "username and password must not be empty")
	}
	user, err := s.q.GetUserByName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, twirp.NewError(twirp.NotFound, "user does not exist")
		}
		s.l.Error("failed to get user by name", zap.Error(err))
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "failed to get user by name"), err)
	}
	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "hash comparison failed"), err)
	}
	token, err := s.auth.MintToken(user.ID)
	if err != nil {
		s.l.Error("failed to mint token", zap.Error(err))
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "failed to mint token"), err)
	}
	resp := &gen.LoginResp{
		Token: token,
	}
	return resp, nil
}
