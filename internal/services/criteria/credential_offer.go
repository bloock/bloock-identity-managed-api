package criteria

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/services/criteria/response"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"strings"
)

type CredentialOffer struct {
	credentialRepository repository.CredentialRepository
	publicHost           string
	logger               zerolog.Logger
}

func NewCredentialOffer(cr repository.CredentialRepository, publicHost string, l zerolog.Logger) *CredentialOffer {
	return &CredentialOffer{
		credentialRepository: cr,
		publicHost:           publicHost,
		logger:               l,
	}
}

func (c CredentialOffer) Get(ctx context.Context, credentialId string) (interface{}, error) {
	credentialUUID, err := uuid.Parse(credentialId)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, domain.ErrInvalidUUID
	}

	credential, err := c.credentialRepository.GetCredentialById(ctx, credentialUUID)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/claims/redeem", strings.TrimSuffix(c.publicHost, "/"))
	id, err := uuid.NewUUID()
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, domain.ErrInvalidUUID
	}

	return response.GetCredentialOfferResponse{
		ID:       id.String(),
		ThreadID: id.String(),
		Body: response.GetCredentialOfferBodyResponse{
			ID:          credential.CredentialId.String(),
			Description: credential.SchemaType,
			URL:         url,
		},
		From: credential.IssuerDid,
		To:   credential.HolderDid,
		Typ:  "application/iden3comm-plain-json",
		Type: "https://iden3-communication.io/credentials/1.0/offer",
	}, nil
}
