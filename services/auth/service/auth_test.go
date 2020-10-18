package service

import (
	"context"
	"github.com/ChenHaoHu/easyupload/services/auth/api"
	"testing"
)

func TestLogin(t *testing.T) {
	auth := NewAuth()
	err := auth.Login(context.Background(), &api.LoginRquest{Name: "hcy", Password: "test"}, &api.LoginResponse{})
	if err != nil {
		t.Fatal(err.Error())
	}
}
