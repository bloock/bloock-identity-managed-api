package request

type CredentialRequest struct {
	SchemaId          string
	SchemaType        string
	IssuerDid         string
	HolderDid         string
	CredentialSubject map[string]interface{}
	Expiration        int64
	Version           int32
	Proofs            []string
}
