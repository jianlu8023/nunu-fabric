// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	handler2 "nunu-fabric/internal/organization/handler"
	repository2 "nunu-fabric/internal/organization/repository"
	server2 "nunu-fabric/internal/organization/server"
	service2 "nunu-fabric/internal/organization/service"
	"nunu-fabric/pkg/helper/sid"
	"nunu-fabric/pkg/jwt"
	"nunu-fabric/pkg/log"
)

// Injectors from wire.go:

func newApp(viperViper *viper.Viper, logger *log.Logger) (*server2.Server, func(), error) {
	jwtJWT := jwt.NewJwt(viperViper)
	handlerHandler := handler2.NewHandler(logger)
	sidSid := sid.NewSid()
	serviceService := service2.NewService(logger, sidSid, jwtJWT)
	db := repository2.NewDB(viperViper)
	client := repository2.NewRedis(viperViper)
	repositoryRepository := repository2.NewRepository(db, client, logger)
	userRepository := repository2.NewUserRepository(repositoryRepository)
	userService := service2.NewUserService(serviceService, userRepository)
	userHandler := handler2.NewUserHandler(handlerHandler, userService)
	engine := server2.NewServerHTTP(logger, jwtJWT, userHandler)
	serverServer := server2.NewServer(engine)
	return serverServer, func() {
	}, nil
}

// wire.go:

var HandlerSet = wire.NewSet(handler2.NewHandler, handler2.NewUserHandler)

var ServiceSet = wire.NewSet(service2.NewService, service2.NewUserService)

var RepositorySet = wire.NewSet(repository2.NewDB, repository2.NewRedis, repository2.NewRepository, repository2.NewUserRepository)
