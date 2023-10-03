package domain

import (
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
	circuits "github.com/iden3/go-circuits/v2"
)

type ProofType int32

const (
	IntegrityProofType ProofType = iota
	SparseMtProofType
	SignatureProofType
)

func NewProofType(proof string) (ProofType, error) {
	switch proof {
	case "integrity_proof":
		return IntegrityProofType, nil
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
	case IntegrityProofType:
		return "integrity_proof"
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

func MapToBloockProofTypes(proofs []ProofType) ([]identityV2.ProofType, error) {
	var proofTypes []identityV2.ProofType
	for _, p := range proofs {
		switch p {
		case IntegrityProofType:
			proofTypes = append(proofTypes, identityV2.IntegrityProofType)
		case SparseMtProofType:
			proofTypes = append(proofTypes, identityV2.SparseMtProofType)
		default:
			return nil, ErrInvalidProofType
		}
	}

	return proofTypes, nil
}

func IsProofIncluded(_type ProofType, proofs []string) bool {
	for _, p := range proofs {
		if p == _type.Str() {
			return true
		}
	}
	return false
}
