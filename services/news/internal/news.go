package internal

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services"
	"github.com/toggleglobal/aaronb-technical-test/services/news/models"
	"github.com/twitchtv/twirp"
)

type Service struct {
	q    *models.Queries
	l    *zap.Logger
	user gen.UserService
}

func NewService(l *zap.Logger, db *sql.DB, user gen.UserService) (*Service, error) {
	q := models.New(db)
	return &Service{
		q:    q,
		l:    l,
		user: user,
	}, nil
}

func toArticle(news models.News) *gen.NewsArticle {
	return &gen.NewsArticle{
		Id:        news.ID,
		Title:     news.Title,
		Timestamp: timestamppb.New(news.Timestamp),
		Tags:      news.Tags,
	}
}

func (s *Service) GetNewsArticle(ctx context.Context, req *gen.GetNewsReq) (*gen.GetNewsResp, error) {
	ts := time.Now()
	if req.LastTimestamp.IsValid() {
		ts = req.LastTimestamp.AsTime()
	}
	// use user tags if none are provided
	if len(req.Tags) == 0 {
		id := services.GetUserIDCtx(ctx)
		if id == uuid.Nil {
			return nil, twirp.NewError(twirp.Unauthenticated, "authentication required")
		}
		userReq := &gen.GetUserTagsReq{
			UserId: id.String(),
		}
		userResp, err := s.user.GetUserTags(ctx, userReq)
		if err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
		req.Tags = userResp.Tags
	}
	query := models.ListNewByTagsPagedParams{
		Timestamp: ts,
		Tags:      req.Tags,
	}
	news, err := s.q.ListNewByTagsPaged(ctx, query)
	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "failed to list news by tag"), err)
	}
	articles := make([]*gen.NewsArticle, len(news))
	for x, v := range news {
		articles[x] = toArticle(v)
	}
	resp := &gen.GetNewsResp{
		Articles: articles,
	}
	return resp, nil
}

func (s *Service) CreateNewsArticle(ctx context.Context, req *gen.CreateNewsReq) (*gen.CreateNewsResp, error) {
	if !req.Article.Timestamp.IsValid() {
		return nil, twirp.NewError(twirp.InvalidArgument, fmt.Sprintf("timestamp is invalid: %v", req.Article.Timestamp))
	}
	params := models.CreateNewsParams{
		Title:     req.Article.Title,
		Timestamp: req.Article.Timestamp.AsTime(),
		Tags:      req.Article.Tags,
	}
	id, err := s.q.CreateNews(ctx, params)
	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "failed to create news"), err)
	}
	resp := &gen.CreateNewsResp{
		Id: id,
	}
	return resp, nil
}
