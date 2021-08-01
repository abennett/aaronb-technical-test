package internal

import (
	"context"
	"database/sql"
    "net/http"
	"errors"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services/news/models"
)

type Service struct {
	q *models.Queries
    Users gen.UserService
	l *zap.Logger
}

func NewService(l *zap.Logger) (*Service, error) {
    // setup postgres
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

    // setup user service client
	userAddr, ok := os.LookupEnv("USER_SRV")
	if !ok {
		return nil, errors.New("USER_SRV not available")
	}
    users := gen.NewUserServiceProtobufClient(userAddr, &http.Client{})
	q := models.New(db)
	return &Service{
		q: q,
        Users: users,
		l: l,
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
	s.l.Info("get_news_request", zap.Any("req", req))
	ts := time.Now()
	if req.LastTimestamp.IsValid() {
		ts = req.LastTimestamp.AsTime()
	}
	query := models.ListNewByTagsPagedParams{
		Timestamp: ts,
		Tags:      req.Tags,
	}
	s.l.Info("query", zap.Any("query", query))
	news, err := s.q.ListNewByTagsPaged(ctx, query)
	if err != nil {
		return nil, err
	}
	s.l.Info("news", zap.Any("news", news))
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
		return nil, fmt.Errorf("timestamp is invalid: %v", req.Article.Timestamp)
	}
	params := models.CreateNewsParams{
		Title:     req.Article.Title,
		Timestamp: req.Article.Timestamp.AsTime(),
		Tags:      req.Article.Tags,
	}
	id, err := s.q.CreateNews(ctx, params)
	if err != nil {
		return nil, err
	}
	resp := &gen.CreateNewsResp{
		Id: id,
	}
	return resp, nil
}
