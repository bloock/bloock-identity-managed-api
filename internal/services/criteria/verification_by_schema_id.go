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
	identityRepository repository.IdentityRepository
	publicUrl          string
	issuer             string
	syncMap            *utils.SyncMap
	logger             zerolog.Logger
}

func NewCreateVerification(ctx context.Context, syncMap *utils.SyncMap, l zerolog.Logger) (*CreateVerification, error) {
	issuerDid := pkg.GetIssuerDidFromContext(ctx)
	if issuerDid == "" {
		return &CreateVerification{}, domain.ErrEmptyIssuerDID
	}

	return &CreateVerification{
		identityRepository: identity.NewIdentityRepository(ctx, l),
		publicUrl:          config.Configuration.Api.PublicHost,
		issuer:             issuerDid,
		syncMap:            syncMap,
		logger:             l,
	}, nil
}

func (c CreateVerification) Create(ctx context.Context, verificationJSON []byte) ([]byte, error) {
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

	request := auth.CreateAuthorizationRequest("verification request", c.issuer, callbackUrl)

	randomUUID := uuid.New().String()
	request.ID = randomUUID
	request.ThreadID = randomUUID
	request.Body.Scope = append(request.Body.Scope, zkRequest)

	c.syncMap.Store(strconv.FormatUint(sessionID, 10), request)

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
