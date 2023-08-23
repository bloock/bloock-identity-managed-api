package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/iden3/go-schema-processor/verifiable"
)

type Credential struct {
	CredentialId   uuid.UUID
	AnchorId       int64
	SchemaType     string
	HolderDid      string
	ProofType      []string
	CredentialData json.RawMessage
	SignatureProof json.RawMessage
	IntegrityProof json.RawMessage
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

	if IsProofIncluded(SparseMtProofType, proofFilter) {
		var sparseMtProof verifiable.Iden3SparseMerkleTreeProof
		if string(c.SparseMtProof) != "null" {
			if err := json.Unmarshal(c.SparseMtProof, &sparseMtProof); err != nil {
				return verifiable.W3CCredential{}, err
			}
			proofs = append(proofs, &sparseMtProof)
		}
	}

	if IsProofIncluded(IntegrityProofType, proofFilter) {
		var integrityProof IntegrityProof
		if string(c.IntegrityProof) != "null" {
			if err := json.Unmarshal(c.IntegrityProof, &integrityProof); err != nil {
				return verifiable.W3CCredential{}, err
			}
			proofs = append(proofs, &integrityProof)
		}
	}
	vc.Proof = proofs

	return vc, nil
}
