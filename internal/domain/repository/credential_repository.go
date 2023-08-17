package repository

import (
	"bloock-identity-managed-api/internal/domain"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	core "github.com/iden3/go-iden3-core"
)

type CredentialRepository interface {
	Save(ctx context.Context, c domain.Credential) error

	GetCredentialById(ctx context.Context, id uuid.UUID) (domain.Credential, error)
	GetCredentialByIssuerAndId(ctx context.Context, issuer *core.DID, id uuid.UUID) (domain.Credential, error)
	FindCredentialsByAnchorId(ctx context.Context, anchorId int64) ([]domain.Credential, error)

	UpdateSignatureProof(ctx context.Context, id uuid.UUID, signatureProof json.RawMessage) error
	UpdateIntegrityProof(ctx context.Context, id uuid.UUID, integrityProof json.RawMessage) error
	UpdateSparseMtProof(ctx context.Context, id uuid.UUID, sparseMtProof json.RawMessage) error
}
