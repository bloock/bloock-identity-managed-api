package domain

import "errors"

var (
	ErrInvalidUUID               = errors.New("invalid uuid format")
	ErrInvalidDID                = errors.New("invalid did format")
	ErrInvalidZkpMessage         = errors.New("invalid zkp message")
	ErrInvalidCredentialSender   = errors.New("credential do not relate to their sender")
	ErrCredentialNotFound        = errors.New("credential not found")
	ErrInvalidProofType          = errors.New("invalid proof type provided")
	ErrInvalidMethodProvided     = errors.New("invalid method provided")
	ErrInvalidBlockchainProvided = errors.New("invalid blockchain provided")
	ErrInvalidDataType           = errors.New("invalid data type provided")
)
