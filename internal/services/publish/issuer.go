package publish

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/identity"
	"bloock-identity-managed-api/internal/platform/key"
	"context"
	"github.com/rs/zerolog"
)

type IssuerPublish struct {
	keyRepository      repository.KeyRepository
	identityRepository repository.IdentityRepository
	didMethod          domain.DidMethod
	logger             zerolog.Logger
}

func NewIssuerPublish(ctx context.Context, l zerolog.Logger) (*IssuerPublish, error) {
	issuerKey := pkg.GetIssuerKeyFromContext(ctx)
	if issuerKey == "" {
		return &IssuerPublish{}, domain.ErrEmptyIssuerKey
	}

	didMethod := domain.PolygonID
	if config.Configuration.Api.PolygonTestEnabled {
		didMethod = domain.PolygonIDTest
	}

	return &IssuerPublish{
		keyRepository:      key.NewKeyRepository(ctx, issuerKey, l),
		identityRepository: identity.NewIdentityRepository(ctx, l),
		didMethod:          didMethod,
		logger:             l,
	}, nil
}

func (i IssuerPublish) Publish(ctx context.Context) (string, error) {
	issuerKey, err := i.keyRepository.LoadIssuerKey(ctx)
	if err != nil {
		return "", err
	}

	issuer, err := i.identityRepository.ImportIssuer(ctx, issuerKey, i.didMethod)
	if err != nil {
		return "", err
	}

	return i.identityRepository.ForcePublishIssuerState(ctx, issuer)
}
