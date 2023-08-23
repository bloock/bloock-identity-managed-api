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

func (i IdentityRepository) GetIssuerByKey(ctx context.Context, issuerKey identityV2.IssuerKey, params identityV2.IssuerParams) (string, error) {
	did, err := i.identityClient.GetIssuerByKey(issuerKey, params)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return "", err
	}

	return did, nil
}

func (i IdentityRepository) CreateSchema(ctx context.Context, issuerId string, req request.CreateSchemaRequest) (string, error) {
	builder := i.identityClient.BuildSchema(req.DisplayName, req.SchemaType, req.Version, req.Description, issuerId)
	var err error

	for _, attr := range req.Attributes {
		builder, err = addAttributeToBuilder(builder, attr)
		if err != nil {
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

func (i IdentityRepository) CreateCredential(ctx context.Context, issuerId string, proofs []domain.ProofType, signer authenticity.BjjSigner, req request.CredentialRequest) (identityV2.CredentialReceipt, error) {
	builder := i.identityClient.BuildCredential(req.SchemaId, req.SchemaType, issuerId, req.HolderDid, req.Expiration, req.Version)
	var err error

	for _, attr := range req.CredentialSubject {
		builder, err = addCredentialSubjectToBuilder(builder, attr)
		if err != nil {
			i.logger.Error().Err(err).Msg("")
			return identityV2.CredentialReceipt{}, err
		}
	}
	proofTypes, err := domain.MapToBloockProofTypes(proofs)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return identityV2.CredentialReceipt{}, err
	}

	credentialReceipt, err := builder.WithProofType(proofTypes).WithSigner(signer).Build()
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

func addAttributeToBuilder(builder identityV2.SchemaBuilder, attr request.AttributeSchema) (identityV2.SchemaBuilder, error) {
	switch attr.DataType {
	case "string":
		return builder.AddStringAttribute(attr.Name, attr.Id, attr.Description, attr.Required), nil
	case "integer":
		return builder.AddNumberAttribute(attr.Name, attr.Id, attr.Description, attr.Required), nil
	case "date":
		return builder.AddDateAttribute(attr.Name, attr.Id, attr.Description, attr.Required), nil
	case "datetime":
		return builder.AddDatetimeAttribute(attr.Name, attr.Id, attr.Description, attr.Required), nil
	case "boolean":
		return builder.AddBooleanAttribute(attr.Name, attr.Id, attr.Description, attr.Required), nil
	default:
		return identityV2.SchemaBuilder{}, domain.ErrInvalidDataType
	}
}

func addCredentialSubjectToBuilder(builder identityV2.CredentialBuilder, cs request.CredentialSubject) (identityV2.CredentialBuilder, error) {
	switch cs.DataType {
	case "string":
		return builder.WithStringAttribute(cs.Key, cs.Value.(string)), nil
	case "integer":
		return builder.WithNumberAttribute(cs.Key, int64(cs.Value.(float64))), nil
	case "date":
		return builder.WithDateAttribute(cs.Key, int64(cs.Value.(float64))), nil
	case "datetime":
		return builder.WithDatetimeAttribute(cs.Key, int64(cs.Value.(float64))), nil
	case "boolean":
		return builder.WithBooleanAttribute(cs.Key, cs.Value.(bool)), nil
	default:
		return identityV2.CredentialBuilder{}, domain.ErrInvalidDataType
	}
}
