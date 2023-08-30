//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"nunu-fabric/internal/xjipc.cas.cn/blockchain/organization/repository"
	"nunu-fabric/pkg/log"
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

func newApp(*viper.Viper, *log.Logger) (*Migrate, func(), error) {
	panic(wire.Build(
		RepositorySet,
		NewMigrate,
	))
}
