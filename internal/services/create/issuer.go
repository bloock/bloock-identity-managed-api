package create

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
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
	params := identityV2.NewIssuerParams()
	if didMethod != "" && didBlockchain != "" && didNetwork != "" {
		method, err := domain.NewDidMethod(didMethod)
		if err != nil {
			i.logger.Error().Err(err).Msg("")
			return nil, err
		}
		blockchain, err := domain.NewDidBlockchain(didBlockchain)
		if err != nil {
			i.logger.Error().Err(err).Msg("")
			return nil, err
		}
		network, err := domain.NewDidNetwork(didNetwork)
		if err != nil {
			i.logger.Error().Err(err).Msg("")
			return nil, err
		}
		params.Method = method.ToBloockMethod()
		params.Blockchain = blockchain.ToBloockBlockchain()
		params.NetworkId = network.ToBloockNetwork()
	}

	issuerKey, err := i.keyRepository.LoadBjjKeyIssuer(ctx)
	if err != nil {
		return nil, err
	}

	return i.identityRepository.CreateIssuer(ctx, issuerKey, params)
}
