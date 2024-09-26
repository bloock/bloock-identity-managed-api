package criteria

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/identity"
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

type CreateVerification struct {
	identityRepository  repository.IdentityRepository
	publicUrl           string
	verifierDid         string
	verificationSyncMap *utils.SyncMap
	authSyncMap         *utils.SyncMap
	logger              zerolog.Logger
}

func NewCreateVerification(ctx context.Context, verificationSyncMap, authSyncMap *utils.SyncMap, l zerolog.Logger) *CreateVerification {
	return &CreateVerification{
		identityRepository:  identity.NewIdentityRepository(ctx, l),
		publicUrl:           config.Configuration.Api.PublicHost,
		verifierDid:         config.Configuration.Verifier.Did,
		verificationSyncMap: verificationSyncMap,
		authSyncMap:         authSyncMap,
		logger:              l,
	}
}

func (c CreateVerification) Create(ctx context.Context, verificationJSON []byte) ([]byte, error) {
	authToken := pkg.GetApiKeyFromContext(ctx)
	if authToken == "" {
		err := domain.ErrEmptyApiKey
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}
	var zkRequest protocol.ZeroKnowledgeProofRequest

	if err := json.Unmarshal(verificationJSON, &zkRequest); err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, domain.ErrInvalidVerificationRequest
	}

	sessionID, err := utils.RandInt64()
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}

	callbackUrl, err := buildCallbackUrl(c.publicUrl, sessionID)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}

	request := auth.CreateAuthorizationRequest("verification request", c.verifierDid, callbackUrl)

	randomUUID := uuid.New().String()
	request.ID = randomUUID
	request.ThreadID = randomUUID
	request.Body.Scope = append(request.Body.Scope, zkRequest)

	c.verificationSyncMap.Store(strconv.FormatUint(sessionID, 10), request)
	c.authSyncMap.Store(strconv.FormatUint(sessionID, 10), pkg.GetApiKeyFromContext(ctx))

	requestBytes, err := json.Marshal(request)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}

	return requestBytes, nil
}

func buildCallbackUrl(publicUrl string, sessionID uint64) (string, error) {
	callbackUrl := fmt.Sprintf("%s%s?sessionId=%d", publicUrl, "/v1/verifications/callback", sessionID)

	return callbackUrl, nil
}
