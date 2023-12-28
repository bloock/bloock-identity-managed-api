package repository

import (
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"github.com/bloock/bloock-sdk-go/v2/entity/authenticity"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
)

type IdentityRepository interface {
	CreateIssuer(ctx context.Context, issuerKey identityV2.IdentityKey, params identityV2.DidParams, name, description, image string, publishInterval int64) (string, error)
	GetIssuerByKey(ctx context.Context, issuerKey identityV2.IdentityKey, params identityV2.DidParams) (string, error)
	PublishIssuerState(ctx context.Context, issuerDid string, signer authenticity.Signer) (string, error)

	CreateCredential(ctx context.Context, issuerId string, signer authenticity.Signer, req request.CredentialRequest) (identityV2.CredentialReceipt, error)
	RevokeCredential(ctx context.Context, signer authenticity.Signer, credential identityV2.Credential) error

	GetSchema(ctx context.Context, schemaID string) (identityV2.Schema, error)
}
