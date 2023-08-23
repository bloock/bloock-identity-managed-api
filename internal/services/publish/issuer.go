package publish

import (
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"github.com/rs/zerolog"
)

type IssuerPublish struct {
	keyRepository      repository.KeyRepository
	identityRepository repository.IdentityRepository
	issuer             string
	logger             zerolog.Logger
}

func NewIssuerPublish(kr repository.KeyRepository, ir repository.IdentityRepository, issuer string, l zerolog.Logger) *IssuerPublish {
	return &IssuerPublish{
		keyRepository:      kr,
		identityRepository: ir,
		issuer:             issuer,
		logger:             l,
	}
}

func (i IssuerPublish) Publish(ctx context.Context) (interface{}, error) {
	bjjSigner, err := i.keyRepository.LoadBjjSigner(ctx)
	if err != nil {
		return nil, err
	}

	return i.identityRepository.PublishIssuerState(ctx, i.issuer, bjjSigner)
}
