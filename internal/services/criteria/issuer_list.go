package criteria

import (
	"bloock-identity-managed-api/internal/domain/repository"
	"context"
	"github.com/rs/zerolog"
)

type IssuerList struct {
	identityRepository repository.IdentityRepository
	logger             zerolog.Logger
}

func NewIssuerList(ir repository.IdentityRepository, l zerolog.Logger) *IssuerList {
	return &IssuerList{
		identityRepository: ir,
		logger:             l,
	}
}

func (i IssuerList) Get(ctx context.Context) (interface{}, error) {
	return i.identityRepository.GetIssuerList(ctx)
}
