package domain

import (
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-schema-processor/verifiable"
)

type IntegrityProof struct {
	Leaves []string    `json:"leaves"`
	Nodes  []string    `json:"nodes"`
	Depth  string      `json:"depth"`
	Bitmap string      `json:"bitmap"`
	Anchor ProofAnchor `json:"anchor"`
	Type   string      `json:"type"`
}

type ProofAnchor struct {
	AnchorID int64           `json:"anchor_id"`
	Networks []AnchorNetwork `json:"networks"`
	Root     string          `json:"root"`
	Status   string          `json:"status"`
}

type AnchorNetwork struct {
	Name   string `json:"name"`
	State  string `json:"state"`
	TxHash string `json:"tx_hash"`
}

func (i IntegrityProof) ProofType() verifiable.ProofType {
	return verifiable.ProofType(i.Type)
}

func (i IntegrityProof) GetCoreClaim() (*core.Claim, error) {
	return &core.Claim{}, nil
}
