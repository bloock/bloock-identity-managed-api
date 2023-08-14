package repository

import (
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
)

type IdentityRepository interface {
	CreateCredential(ctx context.Context, req request.CredentialRequest) error
}
