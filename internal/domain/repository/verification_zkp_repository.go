package repository

import (
	"context"
	"github.com/iden3/iden3comm"
	"github.com/iden3/iden3comm/v2/protocol"
)

type VerificationZkpRepository interface {
	DecodeJWZ(ctx context.Context, token string) (*iden3comm.BasicMessage, error)
	VerifyJWZ(ctx context.Context, token string, request protocol.AuthorizationRequestMessage) error
}
