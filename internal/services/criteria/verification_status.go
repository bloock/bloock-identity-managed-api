package criteria

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/platform/utils"
	"context"
	"github.com/rs/zerolog"
)

type VerificationStatus struct {
	syncMap *utils.SyncMap
	logger  zerolog.Logger
}

func NewVerificationStatus(syncMap *utils.SyncMap, l zerolog.Logger) *VerificationStatus {
	return &VerificationStatus{
		syncMap: syncMap,
		logger:  l,
	}
}

func (v VerificationStatus) Get(ctx context.Context, sessionId string) error {
	res := v.syncMap.Load(sessionId)
	if res == nil {
		err := domain.ErrSessionIdNotFound
		v.logger.Error().Err(err).Msg("")
		return err
	}

	isVerified, ok := res.(bool)
	if !ok {
		err := domain.ErrNotVerified
		v.logger.Error().Err(err).Msg("")
		return err
	}
	if !isVerified {
		err := domain.ErrVerificationFailed
		v.logger.Error().Err(err).Msg("")
		return err
	}

	return nil
}
