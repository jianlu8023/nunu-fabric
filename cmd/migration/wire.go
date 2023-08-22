//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	repository2 "nunu-fabric/internal/cn/cas/xjipc/blockchain/repository"
	"nunu-fabric/pkg/log"
)

var RepositorySet = wire.NewSet(
	repository2.NewDB,
	repository2.NewRedis,
	repository2.NewRepository,
	repository2.NewUserRepository,
)

func newApp(*viper.Viper, *log.Logger) (*Migrate, func(), error) {
	panic(wire.Build(
		RepositorySet,
		NewMigrate,
	))
}
