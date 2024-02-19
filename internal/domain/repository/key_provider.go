package repository

import (
	"context"
	"github.com/bloock/bloock-sdk-go/v2/entity/key"
)

type KeyProvider interface {
	GetIssuerKey(ctx context.Context) (key.Key, error)
}
