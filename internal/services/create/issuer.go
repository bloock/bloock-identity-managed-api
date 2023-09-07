package create

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"github.com/rs/zerolog"
)

type Issuer struct {
	keyRepository      repository.KeyRepository
	identityRepository repository.IdentityRepository
	logger             zerolog.Logger
}

func NewIssuer(kr repository.KeyRepository, ir repository.IdentityRepository, l zerolog.Logger) *Issuer {
	return &Issuer{
		keyRepository:      kr,
		identityRepository: ir,
		logger:             l,
	}
}

func (i Issuer) Create(ctx context.Context, didMethod, didBlockchain, didNetwork string) (interface{}, error) {
	params, err := domain.GetIssuerParams(didMethod, didBlockchain, didNetwork)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return nil, err
	}

	issuerKey, err := i.keyRepository.LoadBjjKeyIssuer(ctx)
	if err != nil {
		return nil, err
	}

	issuerDid, err := i.identityRepository.GetIssuerByKey(ctx, issuerKey, params)
	if err != nil {
		return nil, err
	}
	if issuerDid != "" {
		return issuerDid, nil
	}

	return i.identityRepository.CreateIssuer(ctx, issuerKey, params)
}
