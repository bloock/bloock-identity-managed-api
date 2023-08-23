package local

import (
	"context"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
	"github.com/bloock/bloock-sdk-go/v2/entity/key"
)

type LocalKeyProvider struct {
	client     client.KeyClient
	publicKey  string
	privateKey string
}

func NewLocalKeyProvider(publicLey, privateKey string) LocalKeyProvider {

	return LocalKeyProvider{
		client:     client.NewKeyClient(),
		publicKey:  publicLey,
		privateKey: privateKey,
	}
}

func (l LocalKeyProvider) GetBjjIssuerKey(ctx context.Context) (identityV2.IssuerKey, error) {
	localKey, err := l.client.LoadLocalKey(key.Bjj, l.publicKey, &l.privateKey)
	if err != nil {
		return nil, err
	}

	return identityV2.NewBjjIssuerKey(identityV2.IssuerKeyArgs{LocalKey: &localKey}), nil
}
