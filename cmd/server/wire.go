//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	handler2 "nunu-fabric/internal/cn/cas/xjipc/blockchain/handler"
	repository2 "nunu-fabric/internal/cn/cas/xjipc/blockchain/repository"
	server2 "nunu-fabric/internal/cn/cas/xjipc/blockchain/server"
	service2 "nunu-fabric/internal/cn/cas/xjipc/blockchain/service"
	"nunu-fabric/pkg/helper/sid"
	"nunu-fabric/pkg/jwt"
	"nunu-fabric/pkg/log"
)

var HandlerSet = wire.NewSet(
	handler2.NewHandler,
	handler2.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service2.NewService,
	service2.NewUserService,
)

var RepositorySet = wire.NewSet(
	repository2.NewDB,
	repository2.NewRedis,
	repository2.NewRepository,
	repository2.NewUserRepository,
)

func newApp(*viper.Viper, *log.Logger) (*server2.Server, func(), error) {
	panic(wire.Build(
		RepositorySet,
		ServiceSet,
		HandlerSet,
		server2.NewServer,
		server2.NewServerHTTP,
		sid.NewSid,
		jwt.NewJwt,
	))
}
