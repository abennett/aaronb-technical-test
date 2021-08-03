package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	news     = gen.NewNewsServiceProtobufClient("http://localhost:8100", &http.Client{})
	now      = time.Now()
	articles = []*gen.NewsArticle{
		{
			Title:     "Running aka kicking the earth",
			Timestamp: timestamppb.New(now.Add(time.Hour * -24)),
			Tags:      []string{"health"},
		},
		{
			Title:     "Do you even lift?",
			Timestamp: timestamppb.New(now.Add(time.Hour * -24 * 2)),
			Tags:      []string{"health"},
		},
		{
			Title:     "$$$",
			Timestamp: timestamppb.New(now.Add(time.Hour * -24 * 3)),
			Tags:      []string{"business"},
		},
		{
			Title:     "???",
			Timestamp: timestamppb.New(now.Add(time.Hour * -24 * 4)),
			Tags:      []string{"health", "business"},
		},
	}
)

func TestNews(t *testing.T) {
	user, err := testUser("news_", "business")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	ctx, err = services.InsertAuthHeader(ctx, user.Token)
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range articles {
		req := &gen.CreateNewsReq{
			Article: a,
		}
		_, err := news.CreateNewsArticle(ctx, req)
		if err != nil {
			t.Fatalf("failed to create article: %v", err)
		}
	}
	// a request with no tags should use a user's tags
	req := &gen.GetNewsReq{}
	resp, err := news.GetNewsArticle(ctx, req)
	if err != nil {
		t.Fatalf("failed to get news: %v", err)
	}
	for _, a := range resp.Articles {
		if !tagIsIn("business", a.Tags) {
			t.Fatal("all articles should have a business tag")
		}
	}
	req = &gen.GetNewsReq{
		Tags: []string{"health"},
	}
	resp, err = news.GetNewsArticle(ctx, req)
	if err != nil {
		t.Fatalf("failed to get news: %v", err)
	}
	for _, a := range resp.Articles {
		if !tagIsIn("health", a.Tags) {
			t.Fatal("all articles should have a health tag")
		}
	}
}
