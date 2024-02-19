package repository

import (
	"context"
	"github.com/bloock/bloock-sdk-go/v2/entity/key"
)

type KeyRepository interface {
	LoadIssuerKey(ctx context.Context) (key.Key, error)
}
