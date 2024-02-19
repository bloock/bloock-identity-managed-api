package publish

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/identity"
	"bloock-identity-managed-api/internal/platform/key"
	"context"
	identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"
	"github.com/rs/zerolog"
)

type IssuerPublish struct {
	keyRepository      repository.KeyRepository
	identityRepository repository.IdentityRepository
	didType            identityEntity.DidType
	logger             zerolog.Logger
}

func NewIssuerPublish(ctx context.Context, l zerolog.Logger) (*IssuerPublish, error) {
	issuerKey := pkg.GetIssuerKeyFromContext(ctx)
	if issuerKey == "" {
		return &IssuerPublish{}, domain.ErrEmptyIssuerKey
	}
	method := pkg.GetIssuerDidTypeMethodFromContext(ctx)
	blockchain := pkg.GetIssuerDidTypeBlockchainFromContext(ctx)
	network := pkg.GetIssuerDidTypeNetworkFromContext(ctx)

	didType, err := domain.GetDidType(method, blockchain, network)
	if err != nil {
		return &IssuerPublish{}, err
	}

	return &IssuerPublish{
		keyRepository:      key.NewKeyRepository(ctx, issuerKey, l),
		identityRepository: identity.NewIdentityRepository(ctx, l),
		didType:            didType,
		logger:             l,
	}, nil
}

func (i IssuerPublish) Publish(ctx context.Context) (string, error) {
	issuerKey, err := i.keyRepository.LoadIssuerKey(ctx)
	if err != nil {
		return "", err
	}

	issuer, err := i.identityRepository.ImportIssuer(ctx, issuerKey, i.didType)
	if err != nil {
		return "", err
	}

	return i.identityRepository.ForcePublishIssuerState(ctx, issuer)
}
