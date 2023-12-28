package main

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/repository/sql"
	"bloock-identity-managed-api/internal/platform/repository/sql/connection"
	"bloock-identity-managed-api/internal/platform/server"
	"bloock-identity-managed-api/internal/platform/utils"
	"bloock-identity-managed-api/internal/platform/zkp/loaders"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	auth "github.com/iden3/go-iden3-auth/v2"
	verificationLoader "github.com/iden3/go-iden3-auth/v2/loaders"
	"github.com/iden3/go-iden3-auth/v2/pubsignals"
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	ctx := context.Background()

	// Setup issuer
	issuerDid, err := createIssuer(ctx, logger)
	if err != nil {
		panic(err)
	}

	// Setup configuration
	cfg, err := config.InitConfig(logger, issuerDid)
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

	// Setup circuits loaders
	cls := loaders.NewCircuits("./internal/platform/zkp/credentials/circuits")
	verificationKeyLoader := verificationLoader.FSKeyLoader{Dir: "./internal/platform/zkp/credentials/keys"}

	// Setup resolvers
	resolver := utils.NewBloockNodeResolver(cfg.Blockchain.Provider, cfg.Bloock.ApiKey, cfg.Blockchain.SmartContract)
	resolvers := map[string]pubsignals.StateResolver{
		cfg.Blockchain.ResolverPrefix: resolver,
	}

	// Setup verifier
	verifier, err := auth.NewVerifier(verificationKeyLoader, resolvers, auth.WithIPFSGateway("https://ipfs.io"))
	if err != nil {
		panic(err)
	}

	// Setup Sync Map
	syncMap := utils.NewSyncMap(30 * time.Minute)
	syncMap.CleaningBackground(1 * time.Hour)

	// Setup repositories
	cr := sql.NewSQLCredentialRepository(*conn, 5*time.Second, logger)

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Run API server
	go func() {
		defer wg.Done()
		sr, err := server.NewServer(cr, cls, cfg.Bloock.WebhookSecretKey, logger)
		if err != nil {
			panic(err)
		}
		if err = sr.Start(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func createIssuer(ctx context.Context, logger zerolog.Logger) (string, error) {
	var issuerDid string
	var err error

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
		issuerDid, err = createIssuerService.Create(ctxValue, req)
		if err != nil {
			return "", err
		}
	}
	return issuerDid, nil
}
