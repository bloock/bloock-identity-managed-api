package local

import (
	"context"
	"github.com/bloock/bloock-sdk-go/v2/client"
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

func (l LocalKeyProvider) GetIssuerKey(ctx context.Context) (key.Key, error) {
	localKey, err := l.keyClient.LoadLocalKey(key.Bjj, l.privateKey)
	if err != nil {
		return key.Key{}, err
	}

	return key.Key{LocalKey: &localKey}, nil
}
