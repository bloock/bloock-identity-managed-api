package update

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/iden3/go-schema-processor/verifiable"
	"github.com/rs/zerolog"
)

type SparseMtProofUpdate struct {
	credentialRepository repository.CredentialRepository
	identityRepository   repository.IdentityRepository
	logger               zerolog.Logger
}

func NewSparseMtProofUpdate(cr repository.CredentialRepository, ir repository.IdentityRepository, l zerolog.Logger) *SparseMtProofUpdate {
	return &SparseMtProofUpdate{
		credentialRepository: cr,
		identityRepository:   ir,
		logger:               l,
	}
}

func (s SparseMtProofUpdate) Update(ctx context.Context, credentialId string, sparseMtProof verifiable.Iden3SparseMerkleTreeProof) error {
	credentialUUID, err := uuid.Parse(credentialId)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return domain.ErrInvalidUUID
	}

	credential, err := s.credentialRepository.GetCredentialById(ctx, credentialUUID)
	if err != nil {
		return err
	}

	sparseMtProofBytes, err := json.Marshal(sparseMtProof)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	var updatedSparseMtProof json.RawMessage
	if err = json.Unmarshal(sparseMtProofBytes, &updatedSparseMtProof); err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	return s.credentialRepository.UpdateSparseMtProof(ctx, credential.CredentialId, updatedSparseMtProof)
}
