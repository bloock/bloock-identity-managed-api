package request

type CredentialRequest struct {
	SchemaId          string
	SchemaType        string
	HolderDid         string
	CredentialSubject []CredentialSubject
	Expiration        int64
	Version           int32
	Proofs            []string
}

type CredentialSubject struct {
	DataType string
	Key      string
	Value    interface{}
}
