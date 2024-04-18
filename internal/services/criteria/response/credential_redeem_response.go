package response

import "github.com/iden3/go-schema-processor/v2/verifiable"

type RedeemCredentialResponse struct {
	ID       string
	ThreadID string
	Body     verifiable.W3CCredential
	From     string
	To       string
	Typ      string
	Type     string
}
