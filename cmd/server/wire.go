//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"nunu-fabric/internal/xjipc.cas.cn/blockchain/organization/handler"
	"nunu-fabric/internal/xjipc.cas.cn/blockchain/organization/repository"
	"nunu-fabric/internal/xjipc.cas.cn/blockchain/organization/server"
	"nunu-fabric/internal/xjipc.cas.cn/blockchain/organization/service"
	"nunu-fabric/pkg/helper/sid"
	"nunu-fabric/pkg/jwt"
	"nunu-fabric/pkg/log"
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

func newApp(*viper.Viper, *log.Logger) (*server.Server, func(), error) {
	panic(wire.Build(
		RepositorySet,
		ServiceSet,
		HandlerSet,
		server.NewServer,
		server.NewServerHTTP,
		sid.NewSid,
		jwt.NewJwt,
	))
}
