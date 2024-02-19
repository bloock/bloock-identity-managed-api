package repository

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"
	"github.com/bloock/bloock-sdk-go/v2/entity/key"
)

type IdentityRepository interface {
	CreateIssuer(ctx context.Context, issuerKey key.Key, didType identityEntity.DidType, name, description, image string, publishInterval domain.PublishIntervalMinutes) (string, error)
	ImportIssuer(ctx context.Context, issuerKey key.Key, didType identityEntity.DidType) (identityEntity.Issuer, error)
	ForcePublishIssuerState(ctx context.Context, issuer identityEntity.Issuer) (string, error)

	CreateCredential(ctx context.Context, issuer identityEntity.Issuer, req request.CredentialRequest) (identityEntity.CredentialReceipt, error)
	RevokeCredential(ctx context.Context, credential identityEntity.Credential, issuer identityEntity.Issuer) error

	GetSchema(ctx context.Context, schemaID string) (identityEntity.Schema, error)
}
