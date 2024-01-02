package key

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/key/local"
	"bloock-identity-managed-api/internal/platform/key/managed"
	"context"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/entity/authenticity"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
	"github.com/rs/zerolog"
	"regexp"
)

type KeyRepository struct {
	provider repository.KeyProvider
	logger   zerolog.Logger
}

func NewKeyRepository(ctx context.Context, key string, l zerolog.Logger) KeyRepository {
	l.With().Caller().Str("component", "key-repository").Logger()

	c := client.NewBloockClient(pkg.GetApiKeyFromContext(ctx), config.Configuration.Api.PublicHost, nil)

	var keyProvider repository.KeyProvider
	if isUUID(key) {
		keyProvider = managed.NewManagedKeyProvider(key, c)
	} else {
		keyProvider = local.NewLocalKeyProvider(key, c)
	}

	return KeyRepository{
		provider: keyProvider,
		logger:   l,
	}
}

func (k KeyRepository) LoadBjjKeyIssuer(ctx context.Context) (identityV2.IdentityKey, error) {
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

func isUUID(s string) bool {
	uuidPattern := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

	return uuidPattern.MatchString(s)
}
