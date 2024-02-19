package criteria

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/platform/identity"
	keyRepo "bloock-identity-managed-api/internal/platform/key"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"
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

func (c IssuerByKey) Get(ctx context.Context, req request.DidMetadataRequest) (identityEntity.Issuer, error) {
	params, err := domain.GetDidType(req.Method, req.Blockchain, req.Network)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return identityEntity.Issuer{}, err
	}

	issuerKey, err := c.keyRepository.LoadIssuerKey(ctx)
	if err != nil {
		return identityEntity.Issuer{}, err
	}

	issuer, err := c.identityRepository.ImportIssuer(ctx, issuerKey, params)
	if err != nil {
		return identityEntity.Issuer{}, nil
	}

	return issuer, nil
}
