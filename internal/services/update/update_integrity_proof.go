package update

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"encoding/json"
	"github.com/rs/zerolog"
)

type IntegrityProofUpdate struct {
	credentialRepository repository.CredentialRepository
	identityRepository   repository.IdentityRepository
	logger               zerolog.Logger
}

func NewIntegrityProofUpdate(cr repository.CredentialRepository, ir repository.IdentityRepository, l zerolog.Logger) *IntegrityProofUpdate {
	return &IntegrityProofUpdate{
		credentialRepository: cr,
		identityRepository:   ir,
		logger:               l,
	}
}

func (b IntegrityProofUpdate) Update(ctx context.Context, integrityProof domain.IntegrityProof) error {
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
