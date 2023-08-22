package main

import (
	"fmt"
	"nunu-fabric/pkg/config"
	"nunu-fabric/pkg/http"
	"nunu-fabric/pkg/log"
	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	servers, cleanup, err := newApp(conf, logger)
	if err != nil {
		panic(err)
	}
	logger.Info("server start", zap.String("host", "http://localhost:"+conf.GetString("http.port")))

	//servers.
	http.Run(servers.ServerHTTP, fmt.Sprintf(":%d", conf.GetInt("http.port")))
	defer cleanup()

}
