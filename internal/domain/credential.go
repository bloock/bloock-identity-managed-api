package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/iden3/go-schema-processor/verifiable"
)

type Credential struct {
	CredentialId   uuid.UUID
	SchemaType     string
	IssuerDid      string
	HolderDid      string
	ProofType      []string
	CredentialData json.RawMessage
	SignatureProof json.RawMessage
	BloockProof    json.RawMessage
	SparseMtProof  json.RawMessage
}

func (c Credential) ParseToVerifiableCredential(proofFilter []string) (verifiable.W3CCredential, error) {
	var vc verifiable.W3CCredential

	if err := json.Unmarshal(c.CredentialData, &vc); err != nil {
		return verifiable.W3CCredential{}, err
	}

	proofs := make(verifiable.CredentialProofs, 0)

	var signatureProof verifiable.BJJSignatureProof2021
	if string(c.SignatureProof) != "null" {
		if err := json.Unmarshal(c.SignatureProof, &signatureProof); err != nil {
			return verifiable.W3CCredential{}, err
		}
		proofs = append(proofs, &signatureProof)
	}

	if IsProofIncluded(PolygonProof, proofFilter) {
		var sparseMtProof verifiable.Iden3SparseMerkleTreeProof
		if string(c.SparseMtProof) != "null" {
			if err := json.Unmarshal(c.SparseMtProof, &sparseMtProof); err != nil {
				return verifiable.W3CCredential{}, err
			}
			proofs = append(proofs, &sparseMtProof)
		}
	}

	if IsProofIncluded(BloockProof, proofFilter) {
		var bloockProof IntegrityProof
		if string(c.BloockProof) != "null" {
			if err := json.Unmarshal(c.BloockProof, &bloockProof); err != nil {
				return verifiable.W3CCredential{}, err
			}
			proofs = append(proofs, &bloockProof)
		}
	}
	vc.Proof = proofs

	return vc, nil
}
