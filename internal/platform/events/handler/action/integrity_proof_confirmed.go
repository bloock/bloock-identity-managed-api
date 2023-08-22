package action

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/update"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/rs/zerolog"
)

type IntegrityProofConfirmedEvent struct {
	AnchorID int64  `json:"anchor_id"`
	Proof    string `json:"proof"`
}

type IntegrityProofConfirmed struct {
	updateIntegrityProofService update.IntegrityProofUpdate
	logger                      zerolog.Logger
}

func NewIntegrityProofConfirmed(ip update.IntegrityProofUpdate, l zerolog.Logger) IntegrityProofConfirmed {

	return IntegrityProofConfirmed{
		updateIntegrityProofService: ip,
		logger:                      l,
	}
}

func (i IntegrityProofConfirmed) EventType() string {
	return "identity.integrity_proof_confirmed"
}

func (i IntegrityProofConfirmed) Run(ctx context.Context, event BloockEvent) error {
	dataEventBytes, err := json.Marshal(event.Data)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return err
	}

	var dataEvent IntegrityProofConfirmedEvent
	if err = json.Unmarshal(dataEventBytes, &dataEvent); err != nil {
		i.logger.Error().Err(err).Msg("")
		return err
	}

	proofBytes, err := base64.URLEncoding.DecodeString(dataEvent.Proof)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return err
	}

	var integrityProof domain.IntegrityProof
	if err = json.Unmarshal(proofBytes, &integrityProof); err != nil {
		i.logger.Error().Err(err).Msg("")
		return err
	}

	return i.updateIntegrityProofService.Update(ctx, integrityProof)
}
