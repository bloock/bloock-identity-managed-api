package zkp

import (
	"bloock-identity-managed-api/internal/platform/web3"
	"bloock-identity-managed-api/internal/platform/zkp/loaders"
	"bloock-identity-managed-api/internal/platform/zkp/verify"
	"context"
	"fmt"
	"github.com/iden3/go-circuits"
	"github.com/iden3/go-jwz"
	"github.com/iden3/iden3comm"
	"github.com/iden3/iden3comm/packers"
	"github.com/rs/zerolog"
)

type VerificationZkpRepository struct {
	packageManager iden3comm.PackageManager
	logger         zerolog.Logger
}

func NewVerificationZkpRepository(ctx context.Context, circuitsLoaderService *loaders.Circuits, l zerolog.Logger) (*VerificationZkpRepository, error) {
	l.With().Caller().Str("component", "verification-zkp-repository").Logger()

	wc, err := web3.NewClientWeb3(ctx, l)
	if err != nil {
		l.Error().Err(err).Msg("")
		return &VerificationZkpRepository{}, err
	}

	stateContract, err := wc.GetAbiState()
	if err != nil {
		l.Error().Err(err).Msg("")
		return &VerificationZkpRepository{}, err
	}

	authV2Set, err := circuitsLoaderService.Load(circuits.AuthV2CircuitID)
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

	return &VerificationZkpRepository{
		packageManager: *packageManager,
		logger:         l,
	}, nil
}

func (z *VerificationZkpRepository) PackageManager() iden3comm.PackageManager {
	return z.packageManager
}
