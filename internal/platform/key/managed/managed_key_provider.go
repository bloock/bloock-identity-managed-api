package managed

import (
	"context"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
)

type ManagedKeyProvider struct {
	client client.KeyClient
	keyID  string
}

func NewManagedKeyProvider(keyID string) ManagedKeyProvider {

	return ManagedKeyProvider{
		client: client.NewKeyClient(),
		keyID:  keyID,
	}
}

func (l ManagedKeyProvider) GetBjjIssuerKey(ctx context.Context) (identityV2.IssuerKey, error) {
	managedKey, err := l.client.LoadManagedKey(l.keyID)
	if err != nil {
		return nil, err
	}

	return identityV2.NewBjjIssuerKey(identityV2.IssuerKeyArgs{ManagedKey: &managedKey}), nil
}
