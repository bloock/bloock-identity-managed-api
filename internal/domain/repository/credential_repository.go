package repository

import (
	"bloock-identity-managed-api/internal/domain"
	"context"
	"encoding/json"
	"github.com/google/uuid"
)

type CredentialRepository interface {
	Save(ctx context.Context, c domain.Credential) error

	GetCredentialById(ctx context.Context, id uuid.UUID) (domain.Credential, error)

	UpdateSignatureProof(ctx context.Context, id uuid.UUID, signatureProof json.RawMessage) error
	UpdateSparseMtProof(ctx context.Context, id uuid.UUID, sparseMtProof json.RawMessage) error
}
