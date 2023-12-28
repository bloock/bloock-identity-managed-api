package local

import (
	"context"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/entity/authenticity"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
	"github.com/bloock/bloock-sdk-go/v2/entity/key"
)

type LocalKeyProvider struct {
	keyClient  client.KeyClient
	privateKey string
}

func NewLocalKeyProvider(privateKey string, client client.BloockClient) LocalKeyProvider {
	return LocalKeyProvider{
		keyClient:  client.KeyClient,
		privateKey: privateKey,
	}
}

func (l LocalKeyProvider) GetBjjIssuerKey(ctx context.Context) (identityV2.IdentityKey, error) {
	localKey, err := l.keyClient.LoadLocalKey(key.Bjj, l.privateKey)
	if err != nil {
		return nil, err
	}

	return identityV2.NewBjjIdentityKey(identityV2.IssuerKeyArgs{LocalKey: &localKey}), nil
}

func (l LocalKeyProvider) GetBjjSigner(ctx context.Context) (authenticity.Signer, error) {
	localKey, err := l.keyClient.LoadLocalKey(key.Bjj, l.privateKey)
	if err != nil {
		return authenticity.Signer{}, err
	}

	return authenticity.NewSignerWithLocalKey(localKey, nil), nil
}
