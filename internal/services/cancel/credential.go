package cancel

import (
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"encoding/json"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
	"github.com/iden3/go-schema-processor/verifiable"
	"github.com/rs/zerolog"
)

type CredentialRevocation struct {
	identityRepository repository.IdentityRepository
	logger             zerolog.Logger
}

func NewCredentialRevocation(ir repository.IdentityRepository, l zerolog.Logger) *CredentialRevocation {
	return &CredentialRevocation{
		identityRepository: ir,
		logger:             l,
	}
}

func (c CredentialRevocation) Revoke(ctx context.Context, cred verifiable.W3CCredential) error {
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

	return c.identityRepository.RevokeCredential(ctx, credential)
}
