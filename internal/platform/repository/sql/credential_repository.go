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

func NewSQLCredentialRepository(connection connection.EntConnection, dbTimeout time.Duration, l zerolog.Logger) *SQLCredentialRepository {
	l.With().Caller().Str("component", "credential-repository").Logger()

	return &SQLCredentialRepository{
		connection: connection,
		dbTimeout:  dbTimeout,
		logger:     l,
	}
}

func (s SQLCredentialRepository) Save(ctx context.Context, c domain.Credential) error {
	certificationCreate := s.connection.DB().Credential.Create().
		SetCredentialID(c.CredentialId).
		SetCredentialType(c.CredentialType).
		SetIssuerDid(c.IssuerDid).
		SetHolderDid(c.HolderDid).
		SetCredentialData(c.CredentialData).
		SetSignatureProof(c.SignatureProof).
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
		CredentialType: cs.CredentialType,
		HolderDid:      cs.HolderDid,
		IssuerDid:      cs.IssuerDid,
		CredentialData: cs.CredentialData,
		SignatureProof: cs.SignatureProof,
		SparseMtProof:  cs.SparseMtProof,
	}, nil
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
