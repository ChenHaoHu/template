package main

import (
	"context"
	"github.com/ChenHaoHu/easyupload/services/auth/api"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	mService := micro.NewService()
	mService.Init()
	authService := api.NewAuthService("", mService.Client())
	resp, err := authService.Login(context.Background(), &api.LoginRquest{Name: "hcy", Password: "test"})
	log.Println(resp, err)
	if err == nil {
		log.Println(resp.Token)
	}
}
