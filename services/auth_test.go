package services

import (
    "testing"
    "encoding/base64"
)

const (
    privEnc = "2m/fkzfEb/Kq2QATiWc7lfbdfJe+23bvi+en1UHss+9Fk3tZAn1ImVF0ModZPVqjvIERYmLJDZWODuk61KkbDA=="
    pubEnc = "RZN7WQJ9SJlRdDKHWT1ao7yBEWJiyQ2Vjg7pOtSpGww="
)

func testAuth() (*Auth, error) {
    priv, err := base64.StdEncoding.DecodeString(privEnc)
    if err != nil {
        return nil, err
    }
    pub, err := base64.StdEncoding.DecodeString(pubEnc)
    if err != nil {
        return nil, err
    }
    auth, err := NewAuth(pub, priv)
    if err != nil {
        return nil, err
    }
    return auth, nil
}

func TestAuth(t *testing.T) {
    auth, err := testAuth()
    if err != nil {
        t.Fatal(err)
    }
    token, err := auth.MintToken()
    if err != nil {
        t.Fatal(err)
    }
    t.Log(token)
}
