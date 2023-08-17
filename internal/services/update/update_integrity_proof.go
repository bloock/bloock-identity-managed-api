package update

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"encoding/json"
	"github.com/rs/zerolog"
)

type BloockIntegrityProofUpdate struct {
	credentialRepository repository.CredentialRepository
	identityRepository   repository.IdentityRepository
	logger               zerolog.Logger
}

func NewIntegrityProofUpdate(cr repository.CredentialRepository, ir repository.IdentityRepository, l zerolog.Logger) *BloockIntegrityProofUpdate {
	return &BloockIntegrityProofUpdate{
		credentialRepository: cr,
		identityRepository:   ir,
		logger:               l,
	}
}

func (b BloockIntegrityProofUpdate) Update(ctx context.Context, proof interface{}) error {
	integrityProof, ok := proof.(domain.IntegrityProof)
	if !ok {
		err := domain.ErrInvalidIntegrityProof
		b.logger.Error().Err(err).Msg("")
		return err
	}
	integrityProofBytes, err := json.Marshal(integrityProof)
	if err != nil {
		b.logger.Error().Err(err).Msg("")
		return err
	}

	credentials, err := b.credentialRepository.FindCredentialsByAnchorId(ctx, integrityProof.Anchor.AnchorID)
	if err != nil {
		return err
	}

	for _, credential := range credentials {
		var updatedIntegrityProof json.RawMessage
		if err = json.Unmarshal(integrityProofBytes, &updatedIntegrityProof); err != nil {
			b.logger.Error().Err(err).Msg("")
			return err
		}

		if err = b.credentialRepository.UpdateIntegrityProof(ctx, credential.CredentialId, updatedIntegrityProof); err != nil {
			return err
		}
	}

	return nil
}
