package criteria

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/identity"
	keyRepo "bloock-identity-managed-api/internal/platform/key"
	"bloock-identity-managed-api/internal/platform/utils"
	"bloock-identity-managed-api/internal/services/criteria/response"
	"context"
	"fmt"
	identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"strings"
)

type CredentialOffer struct {
	credentialRepository repository.CredentialRepository
	identityRepository   repository.IdentityRepository
	keyRepository        repository.KeyRepository
	publicHost           string
	didType              identityEntity.DidType
	authSyncMap          *utils.SyncMap
	logger               zerolog.Logger
}

func NewCredentialOffer(ctx context.Context, cr repository.CredentialRepository, authSyncMap *utils.SyncMap, l zerolog.Logger) (*CredentialOffer, error) {
	issuerKey := pkg.GetIssuerKeyFromContext(ctx)
	if issuerKey == "" {
		return &CredentialOffer{}, domain.ErrEmptyIssuerKey
	}
	method := pkg.GetIssuerDidTypeMethodFromContext(ctx)
	blockchain := pkg.GetIssuerDidTypeBlockchainFromContext(ctx)
	network := pkg.GetIssuerDidTypeNetworkFromContext(ctx)

	didType, err := domain.GetDidType(method, blockchain, network)
	if err != nil {
		return &CredentialOffer{}, err
	}

	return &CredentialOffer{
		identityRepository:   identity.NewIdentityRepository(ctx, l),
		keyRepository:        keyRepo.NewKeyRepository(ctx, issuerKey, l),
		credentialRepository: cr,
		publicHost:           config.Configuration.Api.PublicHost,
		didType:              didType,
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

	issuerKey, err := c.keyRepository.LoadIssuerKey(ctx)
	if err != nil {
		return response.GetCredentialOfferResponse{}, err
	}

	issuer, err := c.identityRepository.ImportIssuer(ctx, issuerKey, c.didType)
	if err != nil {
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
		From: issuer.Did.Did,
		To:   credential.HolderDid,
		Typ:  "application/iden3comm-plain-json",
		Type: "https://iden3-communication.io/credentials/1.0/offer",
	}, nil
}
