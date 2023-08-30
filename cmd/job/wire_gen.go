// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"nunu-fabric/internal/organization/job"
	"nunu-fabric/pkg/log"
)

// Injectors from wire.go:

func newApp(viperViper *viper.Viper, logger *log.Logger) (*job.Job, func(), error) {
	jobJob := job.NewJob(logger)
	return jobJob, func() {
	}, nil
}

// wire.go:

var JobSet = wire.NewSet(job.NewJob)
