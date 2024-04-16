package key

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/key/local"
	"bloock-identity-managed-api/internal/platform/key/managed"
	"context"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/entity/key"
	"github.com/rs/zerolog"
	"regexp"
)

type KeyRepository struct {
	provider repository.KeyProvider
	logger   zerolog.Logger
}

func NewKeyRepository(ctx context.Context, key string, l zerolog.Logger) KeyRepository {
	l.With().Caller().Str("component", "key-repository").Logger()

	c := client.NewBloockClient(pkg.GetApiKeyFromContext(ctx), &config.Configuration.Api.PublicHost)

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

func (k KeyRepository) LoadIssuerKey(ctx context.Context) (key.Key, error) {
	issuerKey, err := k.provider.GetIssuerKey(ctx)
	if err != nil {
		k.logger.Error().Err(err).Msg("")
		return key.Key{}, err
	}

	return issuerKey, nil
}

func isUUID(s string) bool {
	uuidPattern := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

	return uuidPattern.MatchString(s)
}
