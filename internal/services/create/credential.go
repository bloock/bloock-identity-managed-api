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
	keyRepository        repository.KeyRepository
	identityRepository   repository.IdentityRepository
	issuer               string
	logger               zerolog.Logger
}

func NewCredential(cr repository.CredentialRepository, ir repository.IdentityRepository, kr repository.KeyRepository, issuer string, l zerolog.Logger) *Credential {
	return &Credential{
		credentialRepository: cr,
		identityRepository:   ir,
		keyRepository:        kr,
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
	proofs := make([]domain.ProofType, 0)
	for _, pr := range req.Proofs {
		proof, err := domain.NewProofType(pr)
		if err != nil {
			c.logger.Error().Err(err).Msg("")
			return nil, err
		}
		proofs = append(proofs, proof)
	}
	signer, err := c.keyRepository.LoadBjjSigner(ctx)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}

	credentialReceipt, err := c.identityRepository.CreateCredential(ctx, c.issuer, proofs, signer, req)
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
