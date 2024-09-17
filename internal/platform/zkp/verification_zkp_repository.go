package zkp

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/platform/web3"
	"bloock-identity-managed-api/internal/platform/zkp/loaders"
	"bloock-identity-managed-api/internal/platform/zkp/verify"
	"context"
	"fmt"
	"github.com/iden3/go-circuits/v2"
	auth "github.com/iden3/go-iden3-auth/v2"
	verificationLoader "github.com/iden3/go-iden3-auth/v2/loaders"
	"github.com/iden3/go-iden3-auth/v2/pubsignals"
	"github.com/iden3/go-jwz/v2"
	"github.com/iden3/iden3comm/v2"
	"github.com/iden3/iden3comm/v2/packers"
	"github.com/iden3/iden3comm/v2/protocol"
	"github.com/rs/zerolog"
	"time"
)

type VerificationZkpRepository struct {
	packageManager iden3comm.PackageManager
	authVerifier   *auth.Verifier
	logger         zerolog.Logger
}

func NewVerificationZkpRepository(ctx context.Context, l zerolog.Logger) (*VerificationZkpRepository, error) {
	l.With().Caller().Str("component", "verification-zkp-repository").Logger()

	// Package manager setup
	cls := loaders.NewCircuits("./internal/platform/zkp/credentials/circuits")

	web3Client, err := web3.NewClientWeb3(ctx, l)
	if err != nil {
		l.Error().Err(err).Msg("")
		return &VerificationZkpRepository{}, err
	}

	stateContract, err := web3Client.GetAbiState()
	if err != nil {
		l.Error().Err(err).Msg("")
		return &VerificationZkpRepository{}, err
	}

	authV2Set, err := cls.Load(circuits.AuthV2CircuitID)
	if err != nil {
		err = fmt.Errorf("failed upload circuits files: %w", err)
		l.Error().Err(err).Msg("")
		return &VerificationZkpRepository{}, err
	}

	verifications := make(map[jwz.ProvingMethodAlg]packers.VerificationParams)
	verifications[jwz.AuthV2Groth16Alg] = packers.NewVerificationParams(authV2Set.VerificationKey,
		verify.StateVerificationHandler(stateContract))
	provers := make(map[jwz.ProvingMethodAlg]packers.ProvingParams)

	zkpPackerV2 := packers.NewZKPPacker(
		provers,
		verifications,
	)

	packageManager := iden3comm.NewPackageManager()

	err = packageManager.RegisterPackers(zkpPackerV2, &packers.PlainMessagePacker{})
	if err != nil {
		l.Error().Err(err).Msg("")
		return &VerificationZkpRepository{}, err
	}

	// Verifier setup
	verificationKeyLoader := verificationLoader.FSKeyLoader{Dir: "./internal/platform/zkp/credentials/keys"}

	resolvers := map[string]pubsignals.StateResolver{
		config.Configuration.Blockchain.ResolverPrefix: web3Client,
	}

	verifier, err := auth.NewVerifier(verificationKeyLoader, resolvers, auth.WithIPFSGateway("https://ipfs.io"))
	if err != nil {
		l.Error().Err(err).Msg("")
		return &VerificationZkpRepository{}, err
	}

	return &VerificationZkpRepository{
		packageManager: *packageManager,
		authVerifier:   verifier,
		logger:         l,
	}, nil
}

func (z VerificationZkpRepository) DecodeJWZ(ctx context.Context, token string) (*iden3comm.BasicMessage, error) {
	message, err := z.packageManager.UnpackWithType(packers.MediaTypeZKPMessage, []byte(token))
	if err != nil {
		z.logger.Error().Err(err).Msg("")
		return &iden3comm.BasicMessage{}, err
	}

	return message, nil
}

func (z VerificationZkpRepository) VerifyJWZ(ctx context.Context, token string, request protocol.AuthorizationRequestMessage) error {
	_, err := z.authVerifier.FullVerify(ctx, token, request, pubsignals.WithAcceptedStateTransitionDelay(5*time.Minute))
	if err != nil {
		z.logger.Error().Err(err).Msg("")
		return err
	}

	return nil
}
