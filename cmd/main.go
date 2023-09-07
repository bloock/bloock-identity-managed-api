package main

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/platform/identity"
	"bloock-identity-managed-api/internal/platform/key"
	"bloock-identity-managed-api/internal/platform/repository/sql"
	"bloock-identity-managed-api/internal/platform/repository/sql/connection"
	"bloock-identity-managed-api/internal/platform/server"
	"bloock-identity-managed-api/internal/platform/web3"
	"bloock-identity-managed-api/internal/platform/zkp"
	"bloock-identity-managed-api/internal/platform/zkp/loaders"
	"bloock-identity-managed-api/internal/services/cancel"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/criteria"
	"bloock-identity-managed-api/internal/services/publish"
	"bloock-identity-managed-api/internal/services/update"
	"context"
	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	ctx := context.Background()

	// GetBjjIssuerKey ent client
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

	bloock.ApiKey = cfg.APIKey

	// Setup circuits loaders
	cls := loaders.NewCircuits("./internal/platform/zkp/credentials/circuits")

	// Setup Web3Client
	wc, err := web3.NewClientWeb3(config.PolygonProvider, config.PolygonSmartContract, logger)
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
	kr, err := key.NewKeyRepository(cfg.LocalPrivateKey, cfg.LocalPublicKey, cfg.ManagedKeyID, logger)
	if err != nil {
		panic(err)
	}
	ir := identity.NewIdentityRepository(cfg.PublicHost, logger)

	// Create or Retrieve Issuer
	ci := create.NewIssuer(kr, ir, logger)
	res, err := ci.Create(ctx, cfg.IssuerDidMethod, cfg.IssuerDidBlockchain, cfg.IssuerDidNetwork)
	if err != nil {
		panic(err)
	}
	issuerDid := res.(string)

	// Setup registry
	cc := create.NewCredential(cr, ir, kr, issuerDid, logger)
	co := criteria.NewCredentialOffer(cr, cfg.PublicHost, issuerDid, logger)
	rc := criteria.NewCredentialRedeem(cr, vr, logger)
	cbi := criteria.NewCredentialById(cr, logger)
	bpu := update.NewIntegrityProofUpdate(cr, ir, logger)
	smp := update.NewSparseMtProofUpdate(cr, ir, logger)
	crv := cancel.NewCredentialRevocation(ir, logger)
	pi := publish.NewIssuerPublish(kr, ir, issuerDid, logger)
	gi := criteria.NewIssuer(ir, kr, cfg.IssuerDidMethod, cfg.IssuerDidBlockchain, cfg.IssuerDidNetwork, logger)

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Run API server
	go func() {
		defer wg.Done()
		sr, err := server.NewServer(cfg.APIHost, cfg.APIPort, *cc, *co, *rc, *cbi, *bpu, *smp, *gi, *crv, *pi, cfg.WebhookSecretKey, logger, cfg.DebugMode)
		if err != nil {
			panic(err)
		}
		if err = sr.Start(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
