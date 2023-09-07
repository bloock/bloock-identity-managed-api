package request

type CredentialRequest struct {
	SchemaId          string
	HolderDid         string
	CredentialSubject []CredentialSubject
	Expiration        int64
	Version           int32
	Proofs            []string
}

type CredentialSubject struct {
	Key   string
	Value interface{}
}
