package main

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/repository/sql"
	"bloock-identity-managed-api/internal/platform/repository/sql/connection"
	"bloock-identity-managed-api/internal/platform/server"
	"bloock-identity-managed-api/internal/platform/utils"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	ctx := context.Background()

	// Setup configuration
	cfg, err := config.InitConfig(logger)
	if err != nil {
		panic(err)
	}

	// Setup issuer
	err = createIssuer(ctx, logger, cfg)
	if err != nil {
		panic(err)
	}

	// Setting ent connection
	entConnector := connection.NewEntConnector(logger)
	conn, err := connection.NewEntConnection(cfg.Db.ConnectionString, entConnector, logger)
	if err != nil {
		panic(err)
	}
	// Execute ent migrations
	err = conn.Migrate(ctx)
	if err != nil {
		panic(err)
	}

	// Setup Sync Map
	expiration := config.Configuration.Verification.Expiration
	verificationSyncMap := utils.NewSyncMap(time.Duration(expiration) * time.Minute)
	verificationSyncMap.CleaningBackground(time.Duration(expiration) * time.Minute)

	authSyncMap := utils.NewSyncMap(time.Duration(expiration) * time.Minute)
	authSyncMap.CleaningBackground(time.Duration(expiration) * time.Minute)

	// Setup repositories
	cr := sql.NewSQLCredentialRepository(*conn, 5*time.Second, logger)

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Run API server
	go func() {
		defer wg.Done()
		sr, err := server.NewServer(cr, verificationSyncMap, authSyncMap, cfg.Bloock.WebhookSecretKey, logger)
		if err != nil {
			panic(err)
		}
		if err = sr.Start(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func createIssuer(ctx context.Context, logger zerolog.Logger, cfg *config.Config) error {
	if config.Configuration.Issuer.Key.Key != "" {
		ctxValue := context.WithValue(ctx, pkg.ApiKeyContextKey, config.Configuration.Bloock.ApiKey)
		createIssuerService := create.NewIssuer(ctxValue, config.Configuration.Issuer.Key.Key, logger)

		req := request.CreateIssuerRequest{
			Key:             config.Configuration.Issuer.Key.Key,
			Name:            config.Configuration.Issuer.Name,
			Description:     config.Configuration.Issuer.Description,
			Image:           config.Configuration.Issuer.Image,
			PublishInterval: config.Configuration.Issuer.PublishInterval,
		}

		issuerDid, err := createIssuerService.Create(ctxValue, req)
		if err != nil {
			return err
		}
		cfg.Issuer.IssuerDid = issuerDid
	}

	return nil
}
