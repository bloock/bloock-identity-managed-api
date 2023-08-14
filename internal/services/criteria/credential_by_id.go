package criteria

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"github.com/google/uuid"
	core "github.com/iden3/go-iden3-core"
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

func (c CredentialById) Get(ctx context.Context, issuerDid, credentialId string) (interface{}, error) {
	credentialUUID, err := uuid.Parse(credentialId)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, domain.ErrInvalidUUID
	}
	issuer, err := core.ParseDID(issuerDid)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, domain.ErrInvalidDID
	}

	credential, err := c.credentialRepository.GetCredentialByIssuerAndId(ctx, issuer, credentialUUID)
	if err != nil {
		return nil, err
	}

	return domain.ParseToVerifiableCredential(credential), nil
}
