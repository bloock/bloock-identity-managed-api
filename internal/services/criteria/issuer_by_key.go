package criteria

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/platform/identity"
	keyRepo "bloock-identity-managed-api/internal/platform/key"
	"context"
	identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"
	"github.com/rs/zerolog"
)

type IssuerByKey struct {
	identityRepository repository.IdentityRepository
	keyRepository      repository.KeyRepository
	didMethod          domain.DidMethod
	logger             zerolog.Logger
}

func NewIssuerByKey(ctx context.Context, key string, l zerolog.Logger) *IssuerByKey {
	didMethod := domain.PolygonID
	if config.Configuration.Api.PolygonTestEnabled {
		didMethod = domain.PolygonIDTest
	}

	return &IssuerByKey{
		identityRepository: identity.NewIdentityRepository(ctx, l),
		keyRepository:      keyRepo.NewKeyRepository(ctx, key, l),
		didMethod:          didMethod,
		logger:             l,
	}
}

func (c IssuerByKey) Get(ctx context.Context) (identityEntity.Issuer, error) {
	issuerKey, err := c.keyRepository.LoadIssuerKey(ctx)
	if err != nil {
		return identityEntity.Issuer{}, err
	}

	issuer, err := c.identityRepository.ImportIssuer(ctx, issuerKey, c.didMethod)
	if err != nil {
		return identityEntity.Issuer{}, nil
	}

	return issuer, nil
}
