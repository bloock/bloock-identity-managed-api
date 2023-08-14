package repository

import (
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/rs/zerolog"
)

type BloockIdentityRepository struct {
	apikey          string
	integrityClient client.IdentityClient
	log             zerolog.Logger
}

func NewBloockIdentityRepository(apikey string, log zerolog.Logger) *BloockIdentityRepository {
	bloock.ApiKey = apikey
	return &BloockIdentityRepository{
		apikey:          apikey,
		integrityClient: client.NewIdentityClient(),
		log:             log,
	}
}

func (b BloockIdentityRepository) CreateCredential(ctx context.Context, req request.CredentialRequest) error {
	return nil
}
