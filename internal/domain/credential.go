package domain

import "github.com/google/uuid"

type Credential struct {
	CredentialId   uuid.UUID
	SchemaType     string
	IssuerDid      string
	HolderDid      string
	CredentialData map[string]interface{}
	Proofs         map[string]interface{}
}

func ParseToVerifiableCredential(c Credential) map[string]interface{} {
	vc := c.CredentialData
	for key, value := range c.Proofs {
		vc[key] = value
	}

	return vc
}
