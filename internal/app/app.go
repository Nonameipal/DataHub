package app

import (
	"context"
	"net/http"
	"time"

	"DataHub/internal/configs"
	delivery "DataHub/internal/delivery/http"
	"DataHub/internal/infrostucture/database"
	"DataHub/internal/logger"
)

func Run() error {
	log := logger.GetLogger()
	if err := configs.Load(); err != nil {
		return err
	}

	ctx := context.Background()
	gormDB, err := database.InitConnection(ctx)
	if err != nil {
		return err
	}
	defer database.CloseConnection(gormDB)

	router := delivery.InitRoutes(delivery.Handlers{})

	server := &http.Server{
		Addr:         ":" + configs.AppSettings.AppParams.PortRun,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Info().Str("address", server.Addr).Msg("server listening")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
