package repository

import (
	"context"
	"github.com/bloock/bloock-sdk-go/v2/entity/authenticity"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
)

type KeyRepository interface {
	LoadBjjKeyIssuer(ctx context.Context) (identityV2.IdentityKey, error)
	LoadBjjSigner(ctx context.Context) (authenticity.Signer, error)
}
