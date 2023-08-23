package create

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Credential struct {
	credentialRepository repository.CredentialRepository
	identityRepository   repository.IdentityRepository
	issuer               string
	logger               zerolog.Logger
}

func NewCredential(cr repository.CredentialRepository, ir repository.IdentityRepository, issuer string, l zerolog.Logger) *Credential {
	return &Credential{
		credentialRepository: cr,
		identityRepository:   ir,
		issuer:               issuer,
		logger:               l,
	}
}

func (c Credential) Create(ctx context.Context, req request.CredentialRequest) (interface{}, error) {
	for _, p := range req.Proofs {
		_, err := domain.NewProofType(p)
		if err != nil {
			c.logger.Error().Err(err).Msg("")
			return nil, err
		}
	}

	credentialReceipt, err := c.identityRepository.CreateCredential(ctx, c.issuer, req)
	if err != nil {
		return nil, err
	}

	credentialUUID, err := uuid.Parse(credentialReceipt.CredentialId)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}
	credentialString, err := credentialReceipt.Credential.ToJson()
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}
	var credentialData json.RawMessage
	_ = json.Unmarshal([]byte(credentialString), &credentialData)
	var signatureProof json.RawMessage
	_ = json.Unmarshal([]byte(credentialReceipt.Credential.Proof.SignatureProof), &signatureProof)

	credential := domain.Credential{
		CredentialId:   credentialUUID,
		AnchorId:       credentialReceipt.AnchorID,
		HolderDid:      req.HolderDid,
		SchemaType:     req.SchemaType,
		ProofType:      req.Proofs,
		CredentialData: credentialData,
		SignatureProof: signatureProof,
	}

	if err = c.credentialRepository.Save(ctx, credential); err != nil {
		return nil, err
	}

	return credentialUUID.String(), nil
}
