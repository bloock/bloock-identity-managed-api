package criteria

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"github.com/google/uuid"
	"github.com/iden3/go-schema-processor/verifiable"
	"github.com/rs/zerolog"
)

type CredentialById struct {
	credentialRepository repository.CredentialRepository
	logger               zerolog.Logger
}

func NewCredentialById(cr repository.CredentialRepository, l zerolog.Logger) *CredentialById {
	return &CredentialById{
		credentialRepository: cr,
		logger:               l,
	}
}

func (c CredentialById) Get(ctx context.Context, credentialId string) (verifiable.W3CCredential, error) {
	credentialUUID, err := uuid.Parse(credentialId)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return verifiable.W3CCredential{}, domain.ErrInvalidUUID
	}

	credential, err := c.credentialRepository.GetCredentialById(ctx, credentialUUID)
	if err != nil {
		return verifiable.W3CCredential{}, err
	}

	return credential.ParseToVerifiableCredential()
}
