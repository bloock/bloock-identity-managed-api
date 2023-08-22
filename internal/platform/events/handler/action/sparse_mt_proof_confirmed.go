package action

import (
	"bloock-identity-managed-api/internal/services/update"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/iden3/go-schema-processor/verifiable"
	"github.com/rs/zerolog"
)

type SparseMtProofConfirmedEvent struct {
	CredentialID string `json:"credential_id"`
	Proof        string `json:"proof"`
}

type SparseMtProofConfirmed struct {
	updateSparseMtProofService update.SparseMtProofUpdate
	logger                     zerolog.Logger
}

func NewSparseMtProofConfirmed(sp update.SparseMtProofUpdate, l zerolog.Logger) SparseMtProofConfirmed {
	return SparseMtProofConfirmed{
		updateSparseMtProofService: sp,
		logger:                     l,
	}
}

func (s SparseMtProofConfirmed) EventType() string {
	return "identity.sparse_mt_proof_confirmed"
}

func (s SparseMtProofConfirmed) Run(ctx context.Context, event BloockEvent) error {
	dataEventBytes, err := json.Marshal(event.Data)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	var dataEvent SparseMtProofConfirmedEvent
	if err = json.Unmarshal(dataEventBytes, &dataEvent); err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	proofBytes, err := base64.URLEncoding.DecodeString(dataEvent.Proof)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	var sparseMtProof verifiable.Iden3SparseMerkleTreeProof
	if err = json.Unmarshal(proofBytes, &sparseMtProof); err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	return s.updateSparseMtProofService.Update(ctx, dataEvent.CredentialID, sparseMtProof)
}
