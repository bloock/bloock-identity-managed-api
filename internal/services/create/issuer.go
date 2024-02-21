package create

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/platform/identity"
	keyRepo "bloock-identity-managed-api/internal/platform/key"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"github.com/rs/zerolog"
)

type Issuer struct {
	keyRepository      repository.KeyRepository
	identityRepository repository.IdentityRepository
	didMethod          domain.DidMethod
	logger             zerolog.Logger
}

func NewIssuer(ctx context.Context, key string, l zerolog.Logger) *Issuer {
	didMethod := domain.PolygonID
	if config.Configuration.Api.PolygonTestEnabled {
		didMethod = domain.PolygonIDTest
	}
	return &Issuer{
		keyRepository:      keyRepo.NewKeyRepository(ctx, key, l),
		identityRepository: identity.NewIdentityRepository(ctx, l),
		didMethod:          didMethod,
		logger:             l,
	}
}

func (i Issuer) Create(ctx context.Context, req request.CreateIssuerRequest) (string, error) {
	issuerKey, err := i.keyRepository.LoadIssuerKey(ctx)
	if err != nil {
		return "", err
	}

	issuer, err := i.identityRepository.ImportIssuer(ctx, issuerKey, i.didMethod)
	if err != nil {
		return "", err
	}
	if issuer.Did.Did != "" {
		return issuer.Did.Did, nil
	}

	publishInterval, err := domain.NewPublishIntervalMinutes(req.PublishInterval)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return "", err
	}

	return i.identityRepository.CreateIssuer(ctx, issuerKey, i.didMethod, req.Name, req.Description, req.Image, publishInterval)
}
