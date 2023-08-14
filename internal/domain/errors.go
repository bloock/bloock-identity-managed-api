package domain

import "errors"

var (
	ErrInvalidUUID             = errors.New("invalid uuid format")
	ErrInvalidDID              = errors.New("invalid did format")
	ErrInvalidZkpMessage       = errors.New("invalid zkp message")
	ErrInvalidCredentialSender = errors.New("credential do not relate to their sender")
	ErrCredentialNotFound      = errors.New("credential not found")
)
