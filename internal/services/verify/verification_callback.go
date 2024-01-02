package verify

import (
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/platform/utils"
	"bloock-identity-managed-api/internal/platform/zkp"
	"context"
	"github.com/iden3/iden3comm/v2/protocol"
	"github.com/rs/zerolog"
)

type VerificationCallback struct {
	syncMap                *utils.SyncMap
	verificationRepository repository.VerificationZkpRepository
	logger                 zerolog.Logger
}

func NewVerificationCallback(ctx context.Context, syncMap *utils.SyncMap, l zerolog.Logger) (*VerificationCallback, error) {
	vr, err := zkp.NewVerificationZkpRepository(ctx, l)
	if err != nil {
		return &VerificationCallback{}, err
	}

	return &VerificationCallback{
		syncMap:                syncMap,
		verificationRepository: vr,
		logger:                 l,
	}, nil
}

func (v VerificationCallback) Verify(ctx context.Context, token string, sessionId string) error {
	authRequest := v.syncMap.Load(sessionId)

	if err := v.verificationRepository.VerifyJWZ(ctx, token, authRequest.(protocol.AuthorizationRequestMessage)); err != nil {
		return err
	}

	v.syncMap.Delete(sessionId)

	v.syncMap.Store(sessionId, true)

	return nil
}
