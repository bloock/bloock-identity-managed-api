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
	"math"
	"time"
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

func (i IdentityRepository) CreateCredential(ctx context.Context, issuerId string, proofs []domain.ProofType, signer authenticity.BjjSigner, req request.CredentialRequest) (identityV2.CredentialReceipt, error) {
	builder := i.identityClient.BuildCredential(req.SchemaId, issuerId, req.HolderDid, req.Expiration, req.Version)
	var err error

	for _, attr := range req.CredentialSubject {
		builder, err = buildCredentialSubject(builder, attr)
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

func (i IdentityRepository) GetSchema(ctx context.Context, schemaID string) (identityV2.Schema, error) {
	return identityV2.Schema{}, nil
}

func buildCredentialSubject(builder identityV2.CredentialBuilder, cs request.CredentialSubject) (identityV2.CredentialBuilder, error) {
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
		return identityV2.CredentialBuilder{}, domain.ErrInvalidDataType
	}
}

func parseStringType(cs request.CredentialSubject, builder identityV2.CredentialBuilder) identityV2.CredentialBuilder {
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
