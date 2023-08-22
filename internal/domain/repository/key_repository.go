package repository

import "context"

type KeyRepository interface {
	CreateKey(ctx context.Context) error
}
