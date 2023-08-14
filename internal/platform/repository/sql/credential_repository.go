package sql

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/platform/repository/sql/connection"
	"bloock-identity-managed-api/internal/platform/repository/sql/ent/credential"
	"context"
	"github.com/google/uuid"
	core "github.com/iden3/go-iden3-core"
	"github.com/rs/zerolog"
	"strings"
	"time"
)

type SQLCredentialRepository struct {
	connection connection.EntConnection
	dbTimeout  time.Duration
	logger     zerolog.Logger
}

func NewSQLCertificationRepository(connection connection.EntConnection, dbTimeout time.Duration, logger zerolog.Logger) *SQLCredentialRepository {
	return &SQLCredentialRepository{connection: connection, dbTimeout: dbTimeout, logger: logger}
}

func (s SQLCredentialRepository) Save(ctx context.Context, c domain.Credential) error {
	certificationCreate := s.connection.DB().Credential.Create().SetCredentialID(c.CredentialId).SetSchemaType(c.SchemaType).SetIssuerDid(c.IssuerDid).SetHolderDid(c.HolderDid).SetCredentialData(c.CredentialData).SetProofs(c.Proofs)

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
		SchemaType:     cs.SchemaType,
		IssuerDid:      cs.IssuerDid,
		HolderDid:      cs.HolderDid,
		CredentialData: cs.CredentialData,
		Proofs:         cs.Proofs,
	}, nil
}

func (s SQLCredentialRepository) GetCredentialByIssuerAndId(ctx context.Context, issuer *core.DID, id uuid.UUID) (domain.Credential, error) {
	cs, err := s.connection.DB().Credential.Query().
		Where(credential.CredentialID(id), credential.IssuerDid(issuer.String())).First(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = domain.ErrCredentialNotFound
			s.logger.Error().Err(err).Msg("")
			return domain.Credential{}, err
		}
		s.logger.Error().Err(err).Msg("")
		return domain.Credential{}, err
	}

	return domain.Credential{
		CredentialId:   cs.CredentialID,
		SchemaType:     cs.SchemaType,
		IssuerDid:      cs.IssuerDid,
		HolderDid:      cs.HolderDid,
		CredentialData: cs.CredentialData,
		Proofs:         cs.Proofs,
	}, nil
}

func (s SQLCredentialRepository) UpdateCertificationAnchor(ctx context.Context, id uuid.UUID, proofs map[string]interface{}) error {
	if _, err := s.connection.DB().Credential.Update().SetProofs(proofs).
		Where(credential.CredentialID(id)).
		Save(ctx); err != nil {
		s.logger.Error().Err(err).Msg("")
		return err
	}

	return nil
}
