package identity

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"errors"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/entity/authenticity"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
	"github.com/rs/zerolog"
)

type IdentityRepository struct {
	identityClient client.IdentityV2Client
	logger         zerolog.Logger
}

func NewIdentityRepository(publicHost string, log zerolog.Logger) *IdentityRepository {

	return &IdentityRepository{
		identityClient: client.NewIdentityV2Client(publicHost),
		logger:         log,
	}
}

func (i IdentityRepository) CreateIssuer(ctx context.Context, issuerKey identityV2.IssuerKey, params identityV2.IssuerParams) (string, error) {
	did, err := i.identityClient.CreateIssuer(issuerKey, params)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return "", err
	}

	return did, nil
}

func (i IdentityRepository) GetIssuerList(ctx context.Context) ([]string, error) {
	issuers, err := i.identityClient.GetIssuerList()
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return []string{}, err
	}

	return issuers, nil
}

func (i IdentityRepository) CreateSchema(ctx context.Context, issuerId string, req request.CreateSchemaRequest) (string, error) {
	builder := i.identityClient.BuildSchema(req.DisplayName, req.SchemaType, req.Version, req.Description, issuerId)

	for _, attr := range req.Attributes {
		if err := addAttributeToBuilder(&builder, attr); err != nil {
			i.logger.Error().Err(err).Msg("")
			return "", err
		}
	}

	schema, err := builder.Build()
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return "", err
	}

	return schema.Id, nil
}

func (i IdentityRepository) CreateCredential(ctx context.Context, issuerId string, req request.CredentialRequest) (identityV2.CredentialReceipt, error) {
	builder := i.identityClient.BuildCredential(req.SchemaId, req.SchemaType, issuerId, req.HolderDid, req.Expiration, req.Version)

	for _, attr := range req.CredentialSubject {
		if err := addCredentialSubjectToBuilder(&builder, attr); err != nil {
			i.logger.Error().Err(err).Msg("")
			return identityV2.CredentialReceipt{}, err
		}
	}

	credentialReceipt, err := builder.Build()
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return identityV2.CredentialReceipt{}, err
	}

	return credentialReceipt, nil
}

func (i IdentityRepository) RevokeCredential(ctx context.Context, credential identityV2.Credential) error {
	ok, err := i.identityClient.RevokeCredential(credential)
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

func (i IdentityRepository) PublishIssuerState(ctx context.Context, issuerDid string, signer authenticity.BjjSigner) (string, error) {
	stateBuilder := i.identityClient.BuildIssuerSatePublisher(issuerDid)
	receipt, err := stateBuilder.WithSigner(signer).Build()
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return "", err
	}

	return receipt.TxHash, nil
}

func addAttributeToBuilder(builder *identityV2.SchemaBuilder, attr request.AttributeSchema) error {
	switch attr.DataType {
	case "string":
		builder.AddStringAttribute(attr.Name, attr.Id, attr.Description, attr.Required)
	case "integer":
		builder.AddNumberAttribute(attr.Name, attr.Id, attr.Description, attr.Required)
	case "date":
		builder.AddDateAttribute(attr.Name, attr.Id, attr.Description, attr.Required)
	case "datetime":
		builder.AddDatetimeAttribute(attr.Name, attr.Id, attr.Description, attr.Required)
	case "boolean":
		builder.AddBooleanAttribute(attr.Name, attr.Id, attr.Description, attr.Required)
	default:
		return domain.ErrInvalidDataType
	}

	return nil
}

func addCredentialSubjectToBuilder(builder *identityV2.CredentialBuilder, cs request.CredentialSubject) error {
	switch cs.DataType {
	case "string":
		builder.WithStringAttribute(cs.Key, cs.Value.(string))
	case "integer":
		builder.WithNumberAttribute(cs.Key, int64(cs.Value.(float64)))
	case "date":
		builder.WithDateAttribute(cs.Key, int64(cs.Value.(float64)))
	case "datetime":
		builder.WithDatetimeAttribute(cs.Key, int64(cs.Value.(float64)))
	case "boolean":
		builder.WithBooleanAttribute(cs.Key, cs.Value.(bool))
	default:
		return domain.ErrInvalidDataType
	}

	return nil
}
