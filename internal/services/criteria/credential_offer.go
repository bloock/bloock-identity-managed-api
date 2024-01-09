package criteria

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/utils"
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
	authSyncMap          *utils.SyncMap
	logger               zerolog.Logger
}

func NewCredentialOffer(ctx context.Context, cr repository.CredentialRepository, authSyncMap *utils.SyncMap, l zerolog.Logger) (*CredentialOffer, error) {
	issuerDid := pkg.GetIssuerDidFromContext(ctx)
	if issuerDid == "" {
		return &CredentialOffer{}, domain.ErrEmptyIssuerDID
	}

	return &CredentialOffer{
		credentialRepository: cr,
		publicHost:           config.Configuration.Api.PublicHost,
		issuer:               issuerDid,
		authSyncMap:          authSyncMap,
		logger:               l,
	}, nil
}

func (c CredentialOffer) Get(ctx context.Context, credentialId string) (response.GetCredentialOfferResponse, error) {
	authToken := pkg.GetApiKeyFromContext(ctx)
	if authToken == "" {
		err := domain.ErrEmptyApiKey
		c.logger.Error().Err(err).Msg("")
		return response.GetCredentialOfferResponse{}, err
	}

	credentialUUID, err := uuid.Parse(credentialId)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return response.GetCredentialOfferResponse{}, domain.ErrInvalidUUID
	}

	credential, err := c.credentialRepository.GetCredentialById(ctx, credentialUUID)
	if err != nil {
		return response.GetCredentialOfferResponse{}, err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return response.GetCredentialOfferResponse{}, domain.ErrInvalidUUID
	}

	url := fmt.Sprintf("%s/v1/credentials/redeem?thread_id=%s", strings.TrimSuffix(c.publicHost, "/"), id.String())

	c.authSyncMap.Store(id.String(), authToken)

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
