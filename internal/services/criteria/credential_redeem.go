package criteria

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/services/criteria/response"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/iden3comm"
	"github.com/iden3/iden3comm/packers"
	"github.com/iden3/iden3comm/protocol"
	"github.com/rs/zerolog"
)

type CredentialRedeem struct {
	credentialRepository   repository.CredentialRepository
	verificationRepository repository.VerificationZkpRepository
	logger                 zerolog.Logger
}

func NewCredentialRedeem(cr repository.CredentialRepository, vr repository.VerificationZkpRepository, l zerolog.Logger) *CredentialRedeem {
	return &CredentialRedeem{
		credentialRepository:   cr,
		verificationRepository: vr,
		logger:                 l,
	}
}

func (c CredentialRedeem) Redeem(ctx context.Context, body string, proofs []string) (interface{}, error) {
	manager := c.verificationRepository.PackageManager()
	basicMessage, err := manager.UnpackWithType(packers.MediaTypeZKPMessage, []byte(body))
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, domain.ErrInvalidZkpMessage
	}
	issuerDID, subjectDID, err := c.validateBasicMessage(basicMessage)
	if err != nil {
		return nil, err
	}

	fetchRequestBody := &protocol.CredentialFetchRequestMessageBody{}
	err = json.Unmarshal(basicMessage.Body, fetchRequestBody)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, domain.ErrInvalidZkpMessage
	}
	credentialUUID, err := uuid.Parse(fetchRequestBody.ID)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, domain.ErrInvalidUUID
	}

	credential, err := c.credentialRepository.GetCredentialById(ctx, credentialUUID)
	if err != nil {
		return nil, err
	}
	if credential.HolderDid != subjectDID.String() {
		err = domain.ErrInvalidCredentialSender
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}

	vc, err := credential.ParseToVerifiableCredential(proofs)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return nil, domain.ErrInvalidUUID
	}

	return response.RedeemCredentialResponse{
		ID:       id.String(),
		ThreadID: basicMessage.ThreadID,
		Body:     vc,
		From:     issuerDID.String(),
		To:       subjectDID.String(),
		Typ:      string(packers.MediaTypePlainMessage),
		Type:     string(protocol.CredentialIssuanceResponseMessageType),
	}, nil
}

func (c CredentialRedeem) validateBasicMessage(basicMessage *iden3comm.BasicMessage) (*core.DID, *core.DID, error) {
	didTo, err := core.ParseDID(basicMessage.To)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return &core.DID{}, &core.DID{}, domain.ErrInvalidDID
	}
	didFrom, err := core.ParseDID(basicMessage.From)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return &core.DID{}, &core.DID{}, domain.ErrInvalidDID
	}
	_, err = uuid.Parse(basicMessage.ID)
	if err != nil {
		c.logger.Error().Err(err).Msg("")
		return &core.DID{}, &core.DID{}, domain.ErrInvalidUUID
	}
	if basicMessage.Type != protocol.CredentialFetchRequestMessageType && basicMessage.Type != protocol.RevocationStatusRequestMessageType {
		err = domain.ErrInvalidZkpMessage
		c.logger.Error().Err(err).Msg("")
		return &core.DID{}, &core.DID{}, err
	}
	if basicMessage.ID == "" {
		err = domain.ErrInvalidZkpMessage
		c.logger.Error().Err(err).Msg("")
		return &core.DID{}, &core.DID{}, err
	}
	return didTo, didFrom, nil
}
