package create

import (
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"github.com/rs/zerolog"
)

type Schema struct {
	identityRepository repository.IdentityRepository
	issuer             string
	logger             zerolog.Logger
}

func NewSchema(ir repository.IdentityRepository, issuer string, l zerolog.Logger) *Schema {
	return &Schema{
		identityRepository: ir,
		issuer:             issuer,
		logger:             l,
	}
}

func (s Schema) Create(ctx context.Context, req request.CreateSchemaRequest) (interface{}, error) {
	return s.identityRepository.CreateSchema(ctx, s.issuer, req)
}
