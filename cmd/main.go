package main

import (
	"bloock-identity-managed-api/internal/platform/config"
	"bloock-identity-managed-api/internal/platform/repository/sql/connection"
	"context"
	"fmt"
	"github.com/rs/zerolog"
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


	fmt.Println("Hello world")
}
