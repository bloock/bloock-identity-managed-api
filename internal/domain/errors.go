package domain

import "errors"

var (
	ErrInvalidUUID                = errors.New("invalid uuid format")
	ErrInvalidDID                 = errors.New("invalid did format")
	ErrInvalidZkpMessage          = errors.New("invalid zkp message")
	ErrInvalidCredentialSender    = errors.New("credential do not relate to their sender")
	ErrCredentialNotFound         = errors.New("credential not found")
	ErrInvalidProofType           = errors.New("invalid proof type provided")
	ErrInvalidDataType            = errors.New("invalid data type provided")
	ErrSessionIdNotFound          = errors.New("session id not found")
	ErrNotVerified                = errors.New("not verified")
	ErrVerificationFailed         = errors.New("verification failed")
	ErrEmptyIssuerKey             = errors.New("issuer key must be provided")
	ErrInvalidVerificationRequest = errors.New("invalid verification request")
	ErrEmptyApiKey                = errors.New("api key not found")
	ErrInvalidDidMethod           = errors.New("invalid did method")
)
