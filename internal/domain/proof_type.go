package domain

type ProofType int32

const (
	BloockProof ProofType = iota
	PolygonProof
)

func NewProofType(proof string) (ProofType, error) {
	switch proof {
	case "bloock_proof":
		return BloockProof, nil
	case "polygon_proof":
		return PolygonProof, nil
	default:
		return 0, ErrInvalidProofType
	}
}

func (p ProofType) Str() string {
	switch p {
	case BloockProof:
		return "bloock_proof"
	case PolygonProof:
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
