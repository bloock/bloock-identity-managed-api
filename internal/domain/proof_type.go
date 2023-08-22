package domain

type ProofType int32

const (
	IntegrityProofType ProofType = iota
	SparseMtProofType
)

func NewProofType(proof string) (ProofType, error) {
	switch proof {
	case "integrity_proof":
		return IntegrityProofType, nil
	case "sparse_mt_proof":
		return SparseMtProofType, nil
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
	default:
		return ""
	}
}

func IsProofIncluded(_type ProofType, proofs []string) bool {
	for _, p := range proofs {
		if p == _type.Str() {
			return true
		}
	}
	return false
}
