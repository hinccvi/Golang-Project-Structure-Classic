package main

import (
	"context"
	"fmt"
	"hinccvi/go-template/config"
	"hinccvi/go-template/lib/gorm"
	"hinccvi/go-template/lib/redis"
	"hinccvi/go-template/log"
	"hinccvi/go-template/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	//	1. Init yaml config
	config.Init()

	//	2. Init logger
	log.Init(config.Conf.AppConfig.Env)

	//	3. Init gorm
	gorm.Init(config.Conf.AppConfig.Env)

	//	4. Init redis
	redis.Init()

	//	5. Init gin router
	router := routers.Init(config.Conf.AppConfig.Env)

	//	6. Init & Start server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.AppConfig.Port),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic("Fail to listen on port", "port", config.Conf.AppConfig.Port, zap.Error(err))
		}
	}()

	log.Info("Listening on", "port", config.Conf.AppConfig.Port)

	//	7. Gracefully shutdown server with 5 sec delay
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Panic("Server failed to shutdown", zap.Error(err))
	}

	log.Info("Server shut down")
}
