package domain

import (
	circuits "github.com/iden3/go-circuits/v2"
)

type ProofType int32

const (
	SparseMtProofType ProofType = iota
	SignatureProofType
)

func NewProofType(proof string) (ProofType, error) {
	switch proof {
	case "sparse_mt_proof":
		return SparseMtProofType, nil
	case "signature_proof":
		return SignatureProofType, nil
	default:
		return 0, ErrInvalidProofType
	}
}

func (p ProofType) Str() string {
	switch p {
	case SparseMtProofType:
		return "sparse_mt_proof"
	case SignatureProofType:
		return "signature_proof"
	default:
		return ""
	}
}

func (p ProofType) VerificationCircuitProof() (circuits.CircuitID, error) {
	switch p {
	case SparseMtProofType:
		return circuits.AtomicQueryMTPV2CircuitID, nil
	case SignatureProofType:
		return circuits.AtomicQuerySigV2CircuitID, nil
	default:
		return "", ErrInvalidProofType
	}
}
