package cancel

import (
	"context"
	"encoding/json"
	"github.com/iden3/go-schema-processor/verifiable"
	"github.com/rs/zerolog"
)

type CredentialRevocation struct {
	logger zerolog.Logger
}

func NewCredentialRevocation(l zerolog.Logger) *CredentialRevocation {
	return &CredentialRevocation{
		logger: l,
	}
}

func (c CredentialRevocation) Revoke(ctx context.Context, credential verifiable.W3CCredential) (interface{}, error) {
	credentialBytes, err := json.Marshal(credential)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}
	_ = string(credentialBytes)

	//TODO parse credentialString into credential type

	//TODO called sdk revoke credential

	return nil, nil
}
