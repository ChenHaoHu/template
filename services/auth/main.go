package main

import (
	"github.com/ChenHaoHu/easyupload/services/auth/api"
	"github.com/ChenHaoHu/easyupload/services/auth/service"
	"github.com/micro/go-micro/v2"
)

func main() {
	mService := micro.NewService(
	//micro.Name("easyupload.services.auth"),
	)
	mService.Init()
	api.RegisterAuthHandler(mService.Server(), &service.Auth{})
	mService.Run()
}
