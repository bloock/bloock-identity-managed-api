package managed

import (
	"context"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/entity/key"
)

type ManagedKeyProvider struct {
	keyClient client.KeyClient
	keyID     string
}

func NewManagedKeyProvider(keyID string, client client.BloockClient) ManagedKeyProvider {
	return ManagedKeyProvider{
		keyClient: client.KeyClient,
		keyID:     keyID,
	}
}

func (m ManagedKeyProvider) GetIssuerKey(ctx context.Context) (key.Key, error) {
	managedKey, err := m.keyClient.LoadManagedKey(m.keyID)
	if err != nil {
		return key.Key{}, err
	}

	return key.Key{ManagedKey: &managedKey}, nil
}
