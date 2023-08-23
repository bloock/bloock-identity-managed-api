package repository

import (
	"context"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
)

type KeyProvider interface {
	GetBjjIssuerKey(ctx context.Context) (identityV2.IssuerKey, error)
}
