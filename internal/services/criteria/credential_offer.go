package criteria

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
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

func NewCredentialOffer(ctx context.Context, cr repository.CredentialRepository, l zerolog.Logger) (*CredentialOffer, error) {
	issuerDid := pkg.GetIssuerDidFromContext(ctx)
	if issuerDid == "" {
		return &CredentialOffer{}, domain.ErrEmptyIssuerDID
	}

	return &CredentialOffer{
		credentialRepository: cr,
		publicHost:           config.Configuration.Api.PublicHost,
		issuer:               issuerDid,
		logger:               l,
	}, nil
}

func (c CredentialOffer) Get(ctx context.Context, credentialId string) (response.GetCredentialOfferResponse, error) {
	credentialUUID, err := uuid.Parse(credentialId)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return response.GetCredentialOfferResponse{}, domain.ErrInvalidUUID
	}

	credential, err := c.credentialRepository.GetCredentialById(ctx, credentialUUID)
	if err != nil {
		return response.GetCredentialOfferResponse{}, err
	}

	url := fmt.Sprintf("%s/v1/credentials/redeem", strings.TrimSuffix(c.publicHost, "/"))
	id, err := uuid.NewUUID()
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return response.GetCredentialOfferResponse{}, domain.ErrInvalidUUID
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
