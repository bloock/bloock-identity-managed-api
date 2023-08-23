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

func NewManagedKeyProvider(keyID string) ManagedKeyProvider {

	return ManagedKeyProvider{
		keyClient: client.NewKeyClient(),
		keyID:     keyID,
	}
}

func (m ManagedKeyProvider) GetBjjIssuerKey(ctx context.Context) (identityV2.IssuerKey, error) {
	managedKey, err := m.keyClient.LoadManagedKey(m.keyID)
	if err != nil {
		return nil, err
	}

	return identityV2.NewBjjIssuerKey(identityV2.IssuerKeyArgs{ManagedKey: &managedKey}), nil
}

func (m ManagedKeyProvider) GetBjjSigner(ctx context.Context) (authenticity.BjjSigner, error) {
	managedKey, err := m.keyClient.LoadManagedKey(m.keyID)
	if err != nil {
		return authenticity.BjjSigner{}, err
	}

	return authenticity.NewBjjSigner(authenticity.SignerArgs{ManagedKey: &managedKey}), nil
}
