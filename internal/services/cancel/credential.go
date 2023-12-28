package cancel

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/identity"
	"bloock-identity-managed-api/internal/platform/key"
	"context"
	"encoding/json"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
	"github.com/iden3/go-schema-processor/verifiable"
	"github.com/rs/zerolog"
)

type CredentialRevocation struct {
	identityRepository repository.IdentityRepository
	keyRepository      repository.KeyRepository
	logger             zerolog.Logger
}

func NewCredentialRevocation(ctx context.Context, l zerolog.Logger) (*CredentialRevocation, error) {
	issuerKey := pkg.GetIssuerKeyFromContext(ctx)
	if issuerKey == "" {
		return &CredentialRevocation{}, domain.ErrEmptyIssuerKey
	}

	return &CredentialRevocation{
		identityRepository: identity.NewIdentityRepository(ctx, l),
		keyRepository:      key.NewKeyRepository(ctx, issuerKey, l),
		logger:             l,
	}, nil
}

func (c CredentialRevocation) Revoke(ctx context.Context, cred verifiable.W3CCredential) error {
	bjjSigner, err := c.keyRepository.LoadBjjSigner(ctx)
	if err != nil {
		return err
	}

	credentialBytes, err := json.Marshal(cred)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return err
	}

	credential, err := identityV2.NewCredentialFromJson(string(credentialBytes))
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return err
	}

	return c.identityRepository.RevokeCredential(ctx, bjjSigner, credential)
}
