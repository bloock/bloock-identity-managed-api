package verify

import (
	"bloock-identity-managed-api/internal/platform/utils"
	"context"
	auth "github.com/iden3/go-iden3-auth/v2"
	"github.com/iden3/go-iden3-auth/v2/pubsignals"
	"github.com/iden3/iden3comm/v2/protocol"
	"github.com/rs/zerolog"
	"time"
)

type VerificationCallback struct {
	syncMap  *utils.SyncMap
	verifier *auth.Verifier
	logger   zerolog.Logger
}

func NewVerificationCallback(verifier *auth.Verifier, syncMap *utils.SyncMap, l zerolog.Logger) *VerificationCallback {
	return &VerificationCallback{
		syncMap:  syncMap,
		verifier: verifier,
		logger:   l,
	}
}

func (v VerificationCallback) Verify(ctx context.Context, token string, sessionId string) (interface{}, error) {
	authRequest := v.syncMap.Load(sessionId)

	_, err := v.verifier.FullVerify(ctx, token, authRequest.(protocol.AuthorizationRequestMessage), pubsignals.WithAcceptedStateTransitionDelay(5*time.Second))
	if err != nil {
		v.logger.Error().Err(err).Msg("")
		return nil, err
	}

	v.syncMap.Delete(sessionId)

	v.syncMap.Store(sessionId, true)

	return nil, nil
}
