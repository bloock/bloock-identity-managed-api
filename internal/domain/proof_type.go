package domain

type ProofType int32

const (
	IntegrityProofType ProofType = iota
	PolygonProofType
)

func NewProofType(proof string) (ProofType, error) {
	switch proof {
	case "integrity_proof":
		return IntegrityProofType, nil
	case "polygon_proof":
		return PolygonProofType, nil
	default:
		return 0, ErrInvalidProofType
	}
}

func (p ProofType) Str() string {
	switch p {
	case IntegrityProofType:
		return "integrity_proof"
	case PolygonProofType:
		return "polygon_proof"
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
