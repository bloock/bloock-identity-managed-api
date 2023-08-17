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

	UpdateCertificationAnchor(ctx context.Context, id uuid.UUID, signatureProof, bloockProof, sparseMtProof json.RawMessage) error
}
