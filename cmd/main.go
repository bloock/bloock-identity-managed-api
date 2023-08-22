package main

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/platform/key"
	"bloock-identity-managed-api/internal/platform/repository"
	"bloock-identity-managed-api/internal/platform/repository/sql"
	"bloock-identity-managed-api/internal/platform/repository/sql/connection"
	"bloock-identity-managed-api/internal/platform/server"
	"bloock-identity-managed-api/internal/platform/web3"
	"bloock-identity-managed-api/internal/platform/zkp"
	"bloock-identity-managed-api/internal/platform/zkp/loaders"
	"bloock-identity-managed-api/internal/services/cancel"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/criteria"
	"bloock-identity-managed-api/internal/services/update"
	"context"
	"github.com/rs/zerolog"
	"sync"
	"time"
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

	// Setup circuits loaders
	cls := loaders.NewCircuits("./internal/platform/zkp/credentials/circuits")

	// Setup Web3Client
	wc, err := web3.NewClientWeb3(cfg.PolygonProvider, config.PolygonSmartContract, logger)
	if err != nil {
		panic(err)
	}
	sc, err := wc.GetAbiState()
	if err != nil {
		panic(err)
	}

	// Setup Package Manager Zkp
	vr, err := zkp.NewVerificationZkpRepository(ctx, sc, cls, logger)
	if err != nil {
		panic(err)
	}

	// Setup repositories
	cr := sql.NewSQLCertificationRepository(*conn, 5*time.Second, logger)
	ir := repository.NewBloockIdentityRepository(cfg.APIKey, logger)
	kr, err := key.NewKeyRepository(cfg.LocalPrivateKey, cfg.ManagedKeyID, cfg.APIKey, logger)
	if err != nil {
		panic(err)
	}

	// Setup registry
	cc := create.NewCredential(cr, ir, logger)
	co := criteria.NewCredentialOffer(cr, cfg.APIHost, logger)
	rc := criteria.NewCredentialRedeem(cr, vr, logger)
	cbi := criteria.NewCredentialById(cr, logger)
	bpu := update.NewIntegrityProofUpdate(cr, ir, logger)
	smp := update.NewSparseMtProofUpdate(cr, ir, logger)
	ci := create.NewIssuer(kr, logger)
	il := criteria.NewIssuerList(logger)
	cs := create.NewSchema(logger)
	crv := cancel.NewCredentialRevocation(logger)

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Run API server
	go func() {
		defer wg.Done()
		sr, err := server.NewServer(cfg.APIHost, cfg.APIPort, *cc, *co, *rc, *cbi, *bpu, *smp, *ci, *il, *cs, *crv, cfg.WebhookSecretKey, cfg.WebhookEnforceTolerance, logger, cfg.DebugMode)
		if err != nil {
			panic(err)
		}
		if err = sr.Start(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()

}
