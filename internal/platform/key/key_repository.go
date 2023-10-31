package key

import (
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/platform/key/local"
	"bloock-identity-managed-api/internal/platform/key/managed"
	"context"
	"errors"
	"github.com/bloock/bloock-sdk-go/v2/entity/authenticity"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
	"github.com/rs/zerolog"
)

type KeyRepository struct {
	provider repository.KeyProvider
	logger   zerolog.Logger
}

func NewKeyRepository(localPrivateKey, localPublicKey, managedKeyID string, l zerolog.Logger) (KeyRepository, error) {
	var keyProvider repository.KeyProvider
	if localPrivateKey != "" && localPublicKey != "" {
		keyProvider = local.NewLocalKeyProvider(localPublicKey, localPrivateKey)
	} else if managedKeyID != "" {
		keyProvider = managed.NewManagedKeyProvider(managedKeyID)
	} else {
		return KeyRepository{}, errors.New("no local or managed key provided")
	}

	return KeyRepository{
		provider: keyProvider,
		logger:   l,
	}, nil
}

func (k KeyRepository) LoadBjjKeyIssuer(ctx context.Context) (identityV2.IssuerKey, error) {
	issuerKey, err := k.provider.GetBjjIssuerKey(ctx)
	if err != nil {
		k.logger.Error().Err(err).Msg("")
		return nil, err
	}

	return issuerKey, nil
}

func (k KeyRepository) LoadBjjSigner(ctx context.Context) (authenticity.Signer, error) {
	bjjSigner, err := k.provider.GetBjjSigner(ctx)
	if err != nil {
		k.logger.Error().Err(err).Msg("")
		return authenticity.Signer{}, err
	}

	return bjjSigner, nil
}
