package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i-sub135/go-rest-blueprint/source/config"
	"github.com/i-sub135/go-rest-blueprint/source/feature/public/healtcheck"
	"github.com/i-sub135/go-rest-blueprint/source/pkg/db"
	"github.com/i-sub135/go-rest-blueprint/source/pkg/logger"
	"github.com/i-sub135/go-rest-blueprint/source/service"
	"github.com/i-sub135/go-rest-blueprint/source/service/middleware"
)

func main() {

	// initial load config
	if err := config.LoadConfig("config.yaml"); err != nil {
		panic(err)
	}
	cfg := config.GetConfig()

	// initial set logging
	logger.Init(cfg.Log.PrettyConsole)

	// open connection DB
	database, err := db.Init()
	if err != nil {
		logger.Error().Err(err).Msg(err.Error())
		panic(err)
	}

	// initial gin
	gin.SetMode(cfg.App.Mode) // Set mode first
	r := gin.New()
	r.Use(middleware.RequestIDMiddleware())
	r.Use(logger.GinZLogger())
	r.Use(gin.Recovery())

	healthcheck := healtcheck.NewHandler(database)

	r.GET("/health", healthcheck.HealtCheck)

	// Mounting routers
	route_api_v1 := r.Group("/api/v1")
	mounthRoute := service.NewRouters(database)
	mounthRoute.MountRouters(route_api_v1)

	svc := &http.Server{
		Addr:           fmt.Sprintf(":%v", cfg.App.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Info().Str("mode", cfg.App.Mode).Msgf("listening on port %v", cfg.App.Port)
	if err := svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error().Err(err).Msg("server error")
	}

}
