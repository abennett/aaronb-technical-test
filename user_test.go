package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/toggleglobal/aaronb-technical-test/gen"
	"github.com/toggleglobal/aaronb-technical-test/services"
)

var (
	userPub = gen.NewPublicUserServiceProtobufClient("http://localhost:8090", &http.Client{})
	userInt = gen.NewUserServiceProtobufClient("http://localhost:8091", &http.Client{})
)

func tagIsIn(tag string, tags []string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

type userSession struct {
	ID       uuid.UUID
	Name     string
	Password string
	Token    string
}

func testUser(prefix string, tags ...string) (*userSession, error) {
	ctx := context.Background()
	username := prefix + time.Now().Format("15:04:05")
	userReq := &gen.CreateUserReq{
		Username: username,
		Password: prefix + "password",
	}
	createResp, err := userPub.CreateUser(ctx, userReq)
	if err != nil {
		return nil, err
	}
	loginReq := &gen.LoginReq{
		Username: username,
		Password: userReq.Password,
	}
	resp, err := userPub.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}
	if tags != nil {
		ctx, err = services.InsertAuthHeader(ctx, resp.Token)
		if err != nil {
			return nil, err
		}
		for _, t := range tags {
			req := &gen.AddUserTagReq{
				Tag: t,
			}
			_, err := userPub.AddUserTag(ctx, req)
			if err != nil {
				return nil, err
			}
		}
	}
	id, err := uuid.Parse(createResp.UserId)
	if err != nil {
		return nil, err
	}
	return &userSession{
		ID:       id,
		Name:     username,
		Password: userReq.Password,
		Token:    resp.Token,
	}, nil
}

func TestUserCreation(t *testing.T) {
	ctx := context.Background()
	username := "test_username_" + time.Now().Format("15:04:05")
	cases := []struct {
		title     string
		req       *gen.CreateUserReq
		shouldErr bool
	}{
		{
			title: "valid",
			req: &gen.CreateUserReq{
				Username: username,
				Password: "secret here",
			},
			shouldErr: false,
		},
		{
			title: "duplicate of valid",
			req: &gen.CreateUserReq{
				Username: username,
				Password: "secret here",
			},
			shouldErr: true,
		},
		{
			title:     "no fields",
			req:       &gen.CreateUserReq{},
			shouldErr: true,
		},
		{
			title: "no username",
			req: &gen.CreateUserReq{
				Password: "123fksksl",
			},
			shouldErr: true,
		},
		{
			title: "no password",
			req: &gen.CreateUserReq{
				Username: "nope",
			},
			shouldErr: true,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			_, err := userPub.CreateUser(ctx, c.req)
			if err != nil {
				t.Log(err)
				if !c.shouldErr {
					t.Fatalf("%s should have succeeded", c.title)
				}
			}
			if err == nil && c.shouldErr {
				t.Fatalf("%s should err", c.title)
			}
		})
	}
}

func TestUserLogin(t *testing.T) {
	ctx := context.Background()

	username := "login_test_" + time.Now().Format("15:04:05")
	// Create a valid user
	req := &gen.CreateUserReq{
		Username: username,
		Password: "login_pass",
	}
	_, err := userPub.CreateUser(ctx, req)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	cases := []struct {
		title     string
		req       *gen.LoginReq
		shouldErr bool
	}{
		{
			title: "valid user",
			req: &gen.LoginReq{
				Username: req.Username,
				Password: req.Password,
			},
			shouldErr: false,
		},
		{
			title: "invalid user",
			req: &gen.LoginReq{
				Username: "doesn't exist",
				Password: "shouldn't matter",
			},
			shouldErr: true,
		},
		{
			title: "no user",
			req: &gen.LoginReq{
				Username: "",
				Password: "shouldn't matter",
			},
			shouldErr: true,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			_, err := userPub.Login(ctx, c.req)
			if err != nil {
				t.Log(err)
				if !c.shouldErr {
					t.Fatalf("failed to login: %v", err)
				}
			}
			if err == nil && c.shouldErr {
				t.Fatalf("%s should err", c.title)
			}
		})
	}
}

func TestTagMutation(t *testing.T) {
	ctx := context.Background()
	session, err := testUser("add_tag_")
	if err != nil {
		t.Fatal(err)
	}
	ctx, err = services.InsertAuthHeader(ctx, session.Token)
	if err != nil {
		t.Fatal(err)
	}

	// inital added tag
	req := &gen.AddUserTagReq{
		Tag: "health",
	}
	_, err = userPub.AddUserTag(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	tagReq := &gen.GetUserTagsReq{
		UserId: session.ID.String(),
	}
	resp, err := userInt.GetUserTags(ctx, tagReq)
	if len(resp.Tags) != 1 || resp.Tags[0] != "health" {
		t.Fatalf("expected health; received %v", resp.Tags)
	}

	// duplicate add
	req = &gen.AddUserTagReq{
		Tag: "health",
	}
	_, err = userPub.AddUserTag(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	tagReq = &gen.GetUserTagsReq{
		UserId: session.ID.String(),
	}
	resp, err = userInt.GetUserTags(ctx, tagReq)
	if len(resp.Tags) != 1 || resp.Tags[0] != "health" {
		t.Fatalf("expected health; received %v", resp.Tags)
	}

	// new add
	req = &gen.AddUserTagReq{
		Tag: "business",
	}
	_, err = userPub.AddUserTag(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	tagReq = &gen.GetUserTagsReq{
		UserId: session.ID.String(),
	}
	resp, err = userInt.GetUserTags(ctx, tagReq)
	if len(resp.Tags) != 2 || !tagIsIn("business", resp.Tags) {
		t.Fatalf("expected business to be in tags; tags %v", resp.Tags)
	}

	// remote tag
	removeReq := &gen.RemoveUserTagReq{
		Tag: "health",
	}
	_, err = userPub.RemoveUserTag(ctx, removeReq)
	resp, err = userInt.GetUserTags(ctx, tagReq)
	if len(resp.Tags) != 1 || tagIsIn("health", resp.Tags) {
		t.Fatalf("expected health to be removed; tags %v", resp.Tags)
	}
}
