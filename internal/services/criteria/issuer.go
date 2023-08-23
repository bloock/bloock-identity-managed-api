package criteria

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"github.com/rs/zerolog"
)

type Issuer struct {
	keyRepository      repository.KeyRepository
	identityRepository repository.IdentityRepository
	didMethod          string
	didBlockchain      string
	didNetwork         string
	logger             zerolog.Logger
}

func NewIssuer(ir repository.IdentityRepository, kr repository.KeyRepository, dm, db, dn string, l zerolog.Logger) *Issuer {
	return &Issuer{
		identityRepository: ir,
		keyRepository:      kr,
		didMethod:          dm,
		didBlockchain:      db,
		didNetwork:         dn,
		logger:             l,
	}
}

func (i Issuer) Get(ctx context.Context) (interface{}, error) {
	issuerKey, err := i.keyRepository.LoadBjjKeyIssuer(ctx)
	if err != nil {
		return nil, err
	}

	params, err := domain.GetIssuerParams(i.didMethod, i.didBlockchain, i.didNetwork)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return nil, err
	}

	return i.identityRepository.GetIssuerByKey(ctx, issuerKey, params)
}
