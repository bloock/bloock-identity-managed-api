package repository

import (
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"github.com/bloock/bloock-sdk-go/v2/entity/authenticity"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
)

type IdentityRepository interface {
	CreateIssuer(ctx context.Context, issuerKey identityV2.IssuerKey, params identityV2.IssuerParams) (string, error)
	GetIssuerList(ctx context.Context) ([]string, error)
	PublishIssuerState(ctx context.Context, issuerDid string, signer authenticity.BjjSigner) (string, error)

	CreateSchema(ctx context.Context, issuerId string, req request.CreateSchemaRequest) (string, error)

	CreateCredential(ctx context.Context, issuerId string, req request.CredentialRequest) (identityV2.CredentialReceipt, error)
	RevokeCredential(ctx context.Context, credential identityV2.Credential) error
}
