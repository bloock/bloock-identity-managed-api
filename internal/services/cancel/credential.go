package cancel

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/identity"
	"bloock-identity-managed-api/internal/platform/key"
	"context"
	"encoding/json"
	identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"
	"github.com/iden3/go-schema-processor/v2/verifiable"
	"github.com/rs/zerolog"
)

type CredentialRevocation struct {
	identityRepository repository.IdentityRepository
	keyRepository      repository.KeyRepository
	didMethod          domain.DidMethod
	logger             zerolog.Logger
}

func NewCredentialRevocation(ctx context.Context, l zerolog.Logger) (*CredentialRevocation, error) {
	issuerKey := pkg.GetIssuerKeyFromContext(ctx)
	if issuerKey == "" {
		return &CredentialRevocation{}, domain.ErrEmptyIssuerKey
	}

	didMethod := domain.PolygonID
	if config.Configuration.Api.PolygonTestEnabled {
		didMethod = domain.PolygonIDTest
	}

	return &CredentialRevocation{
		identityRepository: identity.NewIdentityRepository(ctx, l),
		keyRepository:      key.NewKeyRepository(ctx, issuerKey, l),
		didMethod:          didMethod,
		logger:             l,
	}, nil
}

func (c CredentialRevocation) Revoke(ctx context.Context, cred verifiable.W3CCredential) error {
	issuerKey, err := c.keyRepository.LoadIssuerKey(ctx)
	if err != nil {
		return err
	}

	issuer, err := c.identityRepository.ImportIssuer(ctx, issuerKey, c.didMethod)
	if err != nil {
		return err
	}

	credentialBytes, err := json.Marshal(cred)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return err
	}

	credential, err := identityEntity.NewCredentialFromJson(string(credentialBytes))
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return err
	}

	return c.identityRepository.RevokeCredential(ctx, credential, issuer)
}
