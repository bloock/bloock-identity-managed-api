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
	issuer               string
	logger               zerolog.Logger
}

func NewCredentialOffer(cr repository.CredentialRepository, publicHost, issuer string, l zerolog.Logger) *CredentialOffer {
	return &CredentialOffer{
		credentialRepository: cr,
		publicHost:           publicHost,
		issuer:               issuer,
		logger:               l,
	}
}

func (c CredentialOffer) Get(ctx context.Context, credentialId string, proofs []string) (interface{}, error) {
	credentialUUID, err := uuid.Parse(credentialId)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, domain.ErrInvalidUUID
	}

	for _, p := range proofs {
		if _, err = domain.NewProofType(p); err != nil {
			c.logger.Error().Err(err).Msg("")
			return nil, err
		}
	}

	credential, err := c.credentialRepository.GetCredentialById(ctx, credentialUUID)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/claims/redeem%s", strings.TrimSuffix(c.publicHost, "/"), getQueryProofs(proofs))
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
			Description: credential.CredentialType,
			URL:         url,
		},
		From: c.issuer,
		To:   credential.HolderDid,
		Typ:  "application/iden3comm-plain-json",
		Type: "https://iden3-communication.io/credentials/1.0/offer",
	}, nil
}

func getQueryProofs(proofs []string) string {
	var uri string
	if len(proofs) == 0 {
		return uri
	}
	queryString := strings.Join(proofs, "&proof=")
	uri = "?proof=" + queryString

	return uri
}
