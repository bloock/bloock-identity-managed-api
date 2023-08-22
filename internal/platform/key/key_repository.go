package key

import (
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/platform/key/local"
	"bloock-identity-managed-api/internal/platform/key/managed"
	"errors"
	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/rs/zerolog"
)

type KeyRepository struct {
	provider repository.KeyProvider
	logger   zerolog.Logger
}

func NewKeyRepository(localPrivateKey, managedKeyID string, apiKey string, l zerolog.Logger) (KeyRepository, error) {
	var keyProvider repository.KeyProvider
	if localPrivateKey != "" {
		keyProvider = local.NewLocalKeyProvider(l)
	} else if managedKeyID != "" {
		keyProvider = managed.NewManagedKeyProvider(l)
	} else {
		return KeyRepository{}, errors.New("no local or managed key provided")
	}

	bloock.ApiKey = apiKey
	return KeyRepository{
		provider: keyProvider,
		logger:   l,
	}, nil
}
