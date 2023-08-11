package main

import (
	"bloock-identity-managed-api/internal/platform/config"
	"bloock-identity-managed-api/internal/platform/repository/sql/connection"
	"bloock-identity-managed-api/internal/platform/server"
	"context"
	"github.com/rs/zerolog"
	"sync"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	logger := zerolog.Logger{}
	ctx := context.Background()

	// Create ent client
	entConnector := connection.NewEntConnector(logger)
	// Setting ent connection
	conn, err := connection.NewEntConnection(cfg.DBConnectionString, entConnector, logger)
	if err != nil {
		panic(err)
	}
	// Execute ent migrations
	err = conn.Migrate(ctx)
	if err != nil {
		panic(err)
	}

	// Setup registry


	wg := sync.WaitGroup{}
	wg.Add(1)

	// Run API server
	go func() {
		defer wg.Done()
		sr, err := server.NewServer(cfg.APIHost, cfg.APIPort, cfg.WebhookSecretKey, cfg.WebhookEnforceTolerance, logger, cfg.DebugMode)
		if err != nil {
			panic(err)
		}
		if err = sr.Start(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()

}
