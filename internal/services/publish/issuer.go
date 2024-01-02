package publish

import (
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
	issuer             string
	logger             zerolog.Logger
}

func NewIssuerPublish(ctx context.Context, l zerolog.Logger) (*IssuerPublish, error) {
	issuerDid := pkg.GetIssuerDidFromContext(ctx)
	if issuerDid == "" {
		return &IssuerPublish{}, domain.ErrEmptyIssuerDID
	}
	issuerKey := pkg.GetIssuerKeyFromContext(ctx)
	if issuerKey == "" {
		return &IssuerPublish{}, domain.ErrEmptyIssuerKey
	}

	return &IssuerPublish{
		keyRepository:      key.NewKeyRepository(ctx, issuerKey, l),
		identityRepository: identity.NewIdentityRepository(ctx, l),
		issuer:             issuerDid,
		logger:             l,
	}, nil
}

func (i IssuerPublish) Publish(ctx context.Context) (string, error) {
	bjjSigner, err := i.keyRepository.LoadBjjSigner(ctx)
	if err != nil {
		return "", err
	}

	return i.identityRepository.PublishIssuerState(ctx, i.issuer, bjjSigner)
}
