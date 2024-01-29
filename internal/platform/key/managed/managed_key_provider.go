package managed

import (
	"context"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/entity/authenticity"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
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

func (m ManagedKeyProvider) GetBjjIssuerKey(ctx context.Context) (identityV2.IdentityKey, error) {
	managedKey, err := m.keyClient.LoadManagedKey(m.keyID)
	if err != nil {
		return nil, err
	}

	return identityV2.NewBjjIdentityKey(identityV2.IssuerKeyArgs{ManagedKey: &managedKey}), nil
}

func (m ManagedKeyProvider) GetBjjSigner(ctx context.Context) (authenticity.Signer, error) {
	managedKey, err := m.keyClient.LoadManagedKey(m.keyID)
	if err != nil {
		return authenticity.Signer{}, err
	}

	return authenticity.NewSignerWithManagedKey(managedKey, nil, nil), nil
}
