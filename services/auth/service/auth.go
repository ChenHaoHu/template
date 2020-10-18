package service

import (
	"context"
	"github.com/ChenHaoHu/easyupload/services/auth/api"
	"log"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

//Login
func (a Auth) Login(ctx context.Context, rquest *api.LoginRquest, response *api.LoginResponse) error {
	log.Println("get some request")
	response.Token = "ssssss"
	return nil
}
