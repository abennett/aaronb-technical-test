package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/toggleglobal/aaronb-technical-test/gen"
)

var userPub = gen.NewPublicUserServiceProtobufClient("http://localhost:8090", &http.Client{})

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
