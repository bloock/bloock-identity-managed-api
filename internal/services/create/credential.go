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
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Credential struct {
	credentialRepository repository.CredentialRepository
	keyRepository        repository.KeyRepository
	identityRepository   repository.IdentityRepository
	issuer               string
	logger               zerolog.Logger
}

func NewCredential(ctx context.Context, cr repository.CredentialRepository, l zerolog.Logger) (*Credential, error) {
	issuerDid := pkg.GetIssuerDidFromContext(ctx)
	if issuerDid == "" {
		return &Credential{}, domain.ErrEmptyIssuerDID
	}
	issuerKey := pkg.GetIssuerKeyFromContext(ctx)
	if issuerKey == "" {
		return &Credential{}, domain.ErrEmptyIssuerKey
	}
	return &Credential{
		credentialRepository: cr,
		identityRepository:   identity.NewIdentityRepository(ctx, l),
		keyRepository:        keyRepo.NewKeyRepository(ctx, issuerKey, l),
		issuer:               issuerDid,
		logger:               l,
	}, nil
}

func (c Credential) Create(ctx context.Context, req request.CredentialRequest) (string, error) {
	signer, err := c.keyRepository.LoadBjjSigner(ctx)
	if err != nil {
		return "", err
	}

	credentialReceipt, err := c.identityRepository.CreateCredential(ctx, c.issuer, signer, req)
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
		IssuerDid:      c.issuer,
		CredentialData: credentialData,
		SignatureProof: signatureProof,
	}

	if err = c.credentialRepository.Save(ctx, credential); err != nil {
		return "", err
	}

	return credentialUUID.String(), nil
}
