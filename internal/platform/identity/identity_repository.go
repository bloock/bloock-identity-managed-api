package identity

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"errors"
	"github.com/bloock/bloock-sdk-go/v2/client"
	identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"
	"github.com/bloock/bloock-sdk-go/v2/entity/key"
	"github.com/rs/zerolog"
	"math"
	"time"
)

type IdentityRepository struct {
	identityClient     client.IdentityClient
	identityCoreClient client.IdentityCoreClient
	logger             zerolog.Logger
}

func NewIdentityRepository(ctx context.Context, l zerolog.Logger) *IdentityRepository {
	l.With().Caller().Str("component", "identity-repository").Logger()

	c := client.NewBloockClient(pkg.GetApiKeyFromContext(ctx), &config.Configuration.Api.PublicHost, nil)

	return &IdentityRepository{
		identityClient:     c.IdentityClient,
		identityCoreClient: c.IdentityCoreClient,
		logger:             l,
	}
}

func (i IdentityRepository) CreateIssuer(ctx context.Context, issuerKey key.Key, didMethod domain.DidMethod, name, description, image string, publishInterval domain.PublishIntervalMinutes) (string, error) {
	issuer, err := i.identityClient.CreateIssuer(issuerKey, publishInterval.Params(), didMethod.GetBloockDidMethod(), name, description, image)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return "", err
	}

	return issuer.Did.Did, nil
}

func (i IdentityRepository) ImportIssuer(ctx context.Context, issuerKey key.Key, didMethod domain.DidMethod) (identityEntity.Issuer, error) {
	issuer, err := i.identityClient.ImportIssuer(issuerKey, didMethod.GetBloockDidMethod())
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return identityEntity.Issuer{}, err
	}

	return issuer, nil
}

func (i IdentityRepository) CreateCredential(ctx context.Context, issuer identityEntity.Issuer, req request.CredentialRequest) (identityEntity.CredentialReceipt, error) {
	builder := i.identityCoreClient.BuildCredential(issuer, req.SchemaId, req.HolderDid, req.Expiration, req.Version)
	var err error

	for _, attr := range req.CredentialSubject {
		builder, err = buildCredentialSubject(builder, attr)
		if err != nil {
			i.logger.Error().Err(err).Msg("")
			return identityEntity.CredentialReceipt{}, err
		}
	}

	credentialReceipt, err := builder.Build()
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return identityEntity.CredentialReceipt{}, err
	}

	return credentialReceipt, nil
}

func (i IdentityRepository) RevokeCredential(ctx context.Context, credential identityEntity.Credential, issuer identityEntity.Issuer) error {
	ok, err := i.identityClient.RevokeCredential(credential, issuer)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return err
	}
	if !ok {
		err = errors.New("revocation was unsuccessful")
		i.logger.Error().Err(err).Msg("")
		return err
	}

	return nil
}

func (i IdentityRepository) ForcePublishIssuerState(ctx context.Context, issuer identityEntity.Issuer) (string, error) {
	receipt, err := i.identityClient.ForcePublishIssuerState(issuer)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return "", err
	}

	return receipt.TxHash, nil
}

func (i IdentityRepository) GetSchema(ctx context.Context, schemaID string) (identityEntity.Schema, error) {
	schema, err := i.identityClient.GetSchema(schemaID)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return identityEntity.Schema{}, err
	}
	return schema, nil
}

func buildCredentialSubject(builder identityEntity.CredentialCoreBuilder, cs request.CredentialSubject) (identityEntity.CredentialCoreBuilder, error) {
	switch cs.Value.(type) {
	case string:
		return parseStringType(cs, builder), nil
	case float64:
		value := cs.Value.(float64)
		if value == math.Trunc(value) {
			return builder.WithIntegerAttribute(cs.Key, int64(value)), nil
		} else {
			return builder.WithDecimalAttribute(cs.Key, cs.Value.(float64)), nil
		}
	case bool:
		return builder.WithBooleanAttribute(cs.Key, cs.Value.(bool)), nil
	default:
		return identityEntity.CredentialCoreBuilder{}, domain.ErrInvalidDataType
	}
}

func parseStringType(cs request.CredentialSubject, builder identityEntity.CredentialCoreBuilder) identityEntity.CredentialCoreBuilder {
	input := cs.Value.(string)
	var parsedTime time.Time
	var err error
	parsedTime, err = time.Parse("2006-01-02", input)
	if err == nil {
		return builder.WithDateAttribute(cs.Key, parsedTime)
	}
	parsedTime, err = time.Parse(time.RFC3339, input)
	if err == nil {
		return builder.WithDatetimeAttribute(cs.Key, parsedTime)
	}

	return builder.WithStringAttribute(cs.Key, input)
}
