package create

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/identity"
	keyRepo "bloock-identity-managed-api/internal/platform/key"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"encoding/json"
	identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Credential struct {
	credentialRepository repository.CredentialRepository
	keyRepository        repository.KeyRepository
	identityRepository   repository.IdentityRepository
	logger               zerolog.Logger
	issuer               identityEntity.Issuer
	didType              identityEntity.DidType
}

func NewCredential(ctx context.Context, cr repository.CredentialRepository, l zerolog.Logger) (*Credential, error) {
	issuerKey := pkg.GetIssuerKeyFromContext(ctx)
	if issuerKey == "" {
		return &Credential{}, domain.ErrEmptyIssuerKey
	}
	method := pkg.GetIssuerDidTypeMethodFromContext(ctx)
	blockchain := pkg.GetIssuerDidTypeBlockchainFromContext(ctx)
	network := pkg.GetIssuerDidTypeNetworkFromContext(ctx)

	didType, err := domain.GetDidType(method, blockchain, network)
	if err != nil {
		return &Credential{}, err
	}

	return &Credential{
		credentialRepository: cr,
		identityRepository:   identity.NewIdentityRepository(ctx, l),
		keyRepository:        keyRepo.NewKeyRepository(ctx, issuerKey, l),
		didType:              didType,
		logger:               l,
	}, nil
}

func (c Credential) Create(ctx context.Context, req request.CredentialRequest) (string, error) {
	issuerKey, err := c.keyRepository.LoadIssuerKey(ctx)
	if err != nil {
		return "", err
	}

	issuer, err := c.identityRepository.ImportIssuer(ctx, issuerKey, c.didType)
	if err != nil {
		return "", err
	}

	credentialReceipt, err := c.identityRepository.CreateCredential(ctx, issuer, req)
	if err != nil {
		return "", err
	}

	credentialUUID, err := uuid.Parse(credentialReceipt.CredentialId)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return "", err
	}

	credentialString, err := credentialReceipt.Credential.ToJson()
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return "", err
	}
	var credentialData json.RawMessage
	if err = json.Unmarshal([]byte(credentialString), &credentialData); err != nil {
		c.logger.Error().Err(err).Msg("")
		return "", err
	}
	var signatureProof json.RawMessage
	if err = json.Unmarshal([]byte(credentialReceipt.Credential.Proof.SignatureProof), &signatureProof); err != nil {
		c.logger.Error().Err(err).Msg("")
		return "", err
	}

	credential := domain.Credential{
		CredentialId:   credentialUUID,
		HolderDid:      req.HolderDid,
		CredentialType: credentialReceipt.CredentialType,
		IssuerDid:      issuer.Did.Did,
		CredentialData: credentialData,
		SignatureProof: signatureProof,
	}

	if err = c.credentialRepository.Save(ctx, credential); err != nil {
		return "", err
	}

	return credentialUUID.String(), nil
}
