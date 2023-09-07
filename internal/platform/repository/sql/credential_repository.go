package sql

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/platform/repository/sql/connection"
	"bloock-identity-managed-api/internal/platform/repository/sql/ent/credential"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"time"
)

type SQLCredentialRepository struct {
	connection connection.EntConnection
	dbTimeout  time.Duration
	logger     zerolog.Logger
}

func NewSQLCertificationRepository(connection connection.EntConnection, dbTimeout time.Duration, logger zerolog.Logger) *SQLCredentialRepository {
	return &SQLCredentialRepository{
		connection: connection,
		dbTimeout:  dbTimeout,
		logger:     logger,
	}
}

func (s SQLCredentialRepository) Save(ctx context.Context, c domain.Credential) error {
	certificationCreate := s.connection.DB().Credential.Create().
		SetCredentialID(c.CredentialId).
		SetAnchorID(c.AnchorId).
		SetCredentialType(c.CredentialType).
		SetHolderDid(c.HolderDid).
		SetProofType(c.ProofType).
		SetCredentialData(c.CredentialData).
		SetSignatureProof(c.SignatureProof).
		SetIntegrityProof(c.IntegrityProof).
		SetSparseMtProof(c.SparseMtProof)

	if _, err := certificationCreate.Save(ctx); err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	return nil
}

func (s SQLCredentialRepository) GetCredentialById(ctx context.Context, id uuid.UUID) (domain.Credential, error) {
	cs, err := s.connection.DB().Credential.Query().
		Where(credential.CredentialID(id)).First(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return domain.Credential{}, err
	}

	return domain.Credential{
		CredentialId:   cs.CredentialID,
		AnchorId:       cs.AnchorID,
		CredentialType: cs.CredentialType,
		HolderDid:      cs.HolderDid,
		ProofType:      cs.ProofType,
		CredentialData: cs.CredentialData,
		SignatureProof: cs.SignatureProof,
		IntegrityProof: cs.IntegrityProof,
		SparseMtProof:  cs.SparseMtProof,
	}, nil
}

func (s SQLCredentialRepository) FindCredentialsByAnchorId(ctx context.Context, anchorId int64) ([]domain.Credential, error) {
	cs, err := s.connection.DB().Credential.Query().
		Where(credential.AnchorID(anchorId)).All(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return []domain.Credential{}, err
	}

	credentials := make([]domain.Credential, 0)
	for _, entCredential := range cs {
		credentials = append(credentials, domain.Credential{
			CredentialId:   entCredential.CredentialID,
			AnchorId:       entCredential.AnchorID,
			CredentialType: entCredential.CredentialType,
			HolderDid:      entCredential.HolderDid,
			ProofType:      entCredential.ProofType,
			CredentialData: entCredential.CredentialData,
			SignatureProof: entCredential.SignatureProof,
			IntegrityProof: entCredential.IntegrityProof,
			SparseMtProof:  entCredential.SparseMtProof,
		})
	}

	return credentials, nil
}

func (s SQLCredentialRepository) UpdateSignatureProof(ctx context.Context, id uuid.UUID, signatureProof json.RawMessage) error {
	if _, err := s.connection.DB().Credential.Update().
		SetSignatureProof(signatureProof).
		Where(credential.CredentialID(id)).
		Save(ctx); err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	return nil
}

func (s SQLCredentialRepository) UpdateIntegrityProof(ctx context.Context, id uuid.UUID, integrityProof json.RawMessage) error {
	if _, err := s.connection.DB().Credential.Update().
		SetIntegrityProof(integrityProof).
		Where(credential.CredentialID(id)).
		Save(ctx); err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	return nil
}

func (s SQLCredentialRepository) UpdateSparseMtProof(ctx context.Context, id uuid.UUID, sparseMtProof json.RawMessage) error {
	if _, err := s.connection.DB().Credential.Update().
		SetSparseMtProof(sparseMtProof).
		Where(credential.CredentialID(id)).
		Save(ctx); err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	return nil
}
