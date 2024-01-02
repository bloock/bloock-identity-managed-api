package criteria

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

type IssuerByKey struct {
	identityRepository repository.IdentityRepository
	keyRepository      repository.KeyRepository
	logger             zerolog.Logger
}

func NewIssuerByKey(ctx context.Context, key string, l zerolog.Logger) *IssuerByKey {
	return &IssuerByKey{
		identityRepository: identity.NewIdentityRepository(ctx, l),
		keyRepository:      keyRepo.NewKeyRepository(ctx, key, l),
		logger:             l,
	}
}

func (c IssuerByKey) Get(ctx context.Context, req request.DidMetadataRequest) (string, error) {
	params, err := domain.GetIssuerParams(req.Method, req.Blockchain, req.Network)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return "", err
	}

	issuerKey, err := c.keyRepository.LoadBjjKeyIssuer(ctx)
	if err != nil {
		return "", err
	}

	issuerDid, err := c.identityRepository.GetIssuerByKey(ctx, issuerKey, params)
	if err != nil {
		return "", nil
	}
	if issuerDid == "" {
		return config.Configuration.Issuer.IssuerDid, nil
	}

	return issuerDid, nil
}
