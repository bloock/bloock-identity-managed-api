package main

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/platform/identity"
	"bloock-identity-managed-api/internal/platform/key"
	"bloock-identity-managed-api/internal/platform/repository/sql"
	"bloock-identity-managed-api/internal/platform/repository/sql/connection"
	"bloock-identity-managed-api/internal/platform/server"
	"bloock-identity-managed-api/internal/platform/utils"
	"bloock-identity-managed-api/internal/platform/web3"
	"bloock-identity-managed-api/internal/platform/zkp"
	"bloock-identity-managed-api/internal/platform/zkp/loaders"
	"bloock-identity-managed-api/internal/services/cancel"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/criteria"
	"bloock-identity-managed-api/internal/services/publish"
	"bloock-identity-managed-api/internal/services/update"
	"bloock-identity-managed-api/internal/services/verify"
	"context"
	"github.com/bloock/bloock-sdk-go/v2"
	auth "github.com/iden3/go-iden3-auth/v2"
	verificationLoader "github.com/iden3/go-iden3-auth/v2/loaders"
	"github.com/iden3/go-iden3-auth/v2/pubsignals"
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
	bloock.ApiHost = "https://api.bloock.dev"

	// Setup circuits loaders
	keyDir := "./internal/platform/zkp/credentials/circuits"
	cls := loaders.NewCircuits(keyDir)
	verificationKeyLoader := verificationLoader.FSKeyLoader{Dir: keyDir}

	// Setup resolvers
	resolver := utils.NewBloockNodeResolver(config.PolygonProvider, cfg.APIKey, config.PolygonSmartContract)
	resolvers := map[string]pubsignals.StateResolver{
		config.ResolverPrefix: resolver,
	}

	// Setup verifier
	verifier, err := auth.NewVerifier(verificationKeyLoader, resolvers, auth.WithIPFSGateway("https://ipfs.io"))
	if err != nil {
		panic(err)
	}

	// Setup Web3Client
	wc, err := web3.NewClientWeb3(ctx, config.PolygonProvider, cfg.APIKey, config.PolygonSmartContract, logger)
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

	// Setup Sync Map
	syncMap := utils.NewSyncMap()

	// Setup repositories
	cr := sql.NewSQLCertificationRepository(*conn, 5*time.Second, logger)
	kr, err := key.NewKeyRepository(cfg.LocalPrivateKey, cfg.LocalPublicKey, cfg.ManagedKeyID, logger)
	if err != nil {
		panic(err)
	}
	ir := identity.NewIdentityRepository(cfg.PublicHost, logger)

	// Create or Retrieve Issuer
	ci := create.NewIssuer(kr, ir, logger)
	res, err := ci.Create(ctx, "", "", "")
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
	gi := criteria.NewIssuer(ir, kr, "", "", "", logger)
	vbs := criteria.NewVerificationBySchemaId(ir, issuerDid, cfg.PublicHost, syncMap, logger)
	vc := verify.NewVerificationCallback(verifier, syncMap, logger)
	vs := criteria.NewVerificationStatus(syncMap, logger)

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Run API server
	go func() {
		defer wg.Done()
		sr, err := server.NewServer(cfg.APIHost, cfg.APIPort, *cc, *co, *rc, *cbi, *bpu, *smp, *gi, *crv, *pi, *vbs, *vc, *vs, cfg.WebhookSecretKey, logger, cfg.DebugMode)
		if err != nil {
			panic(err)
		}
		if err = sr.Start(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
