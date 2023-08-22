// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	repository2 "nunu-fabric/internal/cn/cas/xjipc/blockchain/repository"
	"nunu-fabric/pkg/log"
)

// Injectors from wire.go:

func newApp(viperViper *viper.Viper, logger *log.Logger) (*Migrate, func(), error) {
	db := repository2.NewDB(viperViper)
	migrate := NewMigrate(db, logger)
	return migrate, func() {
	}, nil
}

// wire.go:

var RepositorySet = wire.NewSet(repository2.NewDB, repository2.NewRedis, repository2.NewRepository, repository2.NewUserRepository)