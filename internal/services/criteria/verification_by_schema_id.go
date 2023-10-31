package criteria

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/platform/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	auth "github.com/iden3/go-iden3-auth/v2"
	"github.com/iden3/iden3comm/v2/protocol"
	"github.com/rs/zerolog"
	"strconv"
)

type VerificationBySchemaId struct {
	identityRepository repository.IdentityRepository
	issuerDid          string
	publicUrl          string
	syncMap            *utils.SyncMap
	logger             zerolog.Logger
}

func NewVerificationBySchemaId(ir repository.IdentityRepository, issuerDid, publicUrl string, syncMap *utils.SyncMap, l zerolog.Logger) *VerificationBySchemaId {
	return &VerificationBySchemaId{
		identityRepository: ir,
		issuerDid:          issuerDid,
		publicUrl:          publicUrl,
		syncMap:            syncMap,
		logger:             l,
	}
}

func (c VerificationBySchemaId) Get(ctx context.Context, schemaID string, proof string) (interface{}, error) {
	proofType, err := domain.NewProofType(proof)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}

	schema, err := c.identityRepository.GetSchema(ctx, schemaID)
	if err != nil {
		return nil, err
	}

	sessionID, err := utils.RandInt64()
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return "", err
	}

	callbackUrl, err := buildCallbackUrl(c.publicUrl, sessionID)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}
	proofCircuit, err := proofType.VerificationCircuitProof()
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}

	request := auth.CreateAuthorizationRequest("credential request", c.issuerDid, callbackUrl)

	randomUUID := uuid.New().String()
	request.ID = randomUUID
	request.ThreadID = randomUUID

	var mtpProofRequest protocol.ZeroKnowledgeProofRequest
	mtpProofRequest.ID = utils.RandInt32()
	mtpProofRequest.CircuitID = string(proofCircuit)
	mtpProofRequest.Query = map[string]interface{}{
		"allowedIssuers": []string{c.issuerDid},
		"context":        schema.CidJsonLd,
		"type":           schema.SchemaType,
	}
	request.Body.Scope = append(request.Body.Scope, mtpProofRequest)

	c.syncMap.Store(strconv.FormatUint(sessionID, 10), request)

	requestBytes, err := json.Marshal(request)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}

	return requestBytes, nil
}

func buildCallbackUrl(publicUrl string, sessionID uint64) (string, error) {
	callbackUrl := fmt.Sprintf("%s%s?sessionId=%d", publicUrl, "/v1/verification/callback", sessionID)

	return callbackUrl, nil
}
