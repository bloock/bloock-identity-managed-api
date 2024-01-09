package verify

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/pkg"
	"bloock-identity-managed-api/internal/platform/utils"
	"bloock-identity-managed-api/internal/platform/zkp"
	"context"
	"errors"
	"github.com/iden3/iden3comm/v2/protocol"
	"github.com/rs/zerolog"
)

type VerificationCallback struct {
	verificationSyncMap    *utils.SyncMap
	authSyncMap            *utils.SyncMap
	verificationRepository repository.VerificationZkpRepository
	logger                 zerolog.Logger
}

func NewVerificationCallback(ctx context.Context, verificationSyncMap, authSyncMap *utils.SyncMap, sessionID string, l zerolog.Logger) (*VerificationCallback, error) {
	val := authSyncMap.Load(sessionID)
	token, ok := val.(string)
	if !ok {
		return &VerificationCallback{}, errors.New("session id auth token not found")
	}

	vr, err := zkp.NewVerificationZkpRepository(context.WithValue(ctx, pkg.ApiKeyContextKey, token), l)
	if err != nil {
		return &VerificationCallback{}, err
	}

	return &VerificationCallback{
		verificationSyncMap:    verificationSyncMap,
		authSyncMap:            authSyncMap,
		verificationRepository: vr,
		logger:                 l,
	}, nil
}

func (v VerificationCallback) Verify(ctx context.Context, token string, sessionId string) error {
	val := v.verificationSyncMap.Load(sessionId)
	authRequest, ok := val.(protocol.AuthorizationRequestMessage)
	if !ok {
		return domain.ErrSessionIdNotFound
	}

	if err := v.verificationRepository.VerifyJWZ(ctx, token, authRequest); err != nil {
		return err
	}

	v.verificationSyncMap.Delete(sessionId)

	v.verificationSyncMap.Store(sessionId, true)

	return nil
}
