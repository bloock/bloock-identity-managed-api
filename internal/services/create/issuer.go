package create

import (
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"github.com/rs/zerolog"
)

type Issuer struct {
	keyRepository repository.KeyRepository
	logger        zerolog.Logger
}

func NewIssuer(kr repository.KeyRepository, l zerolog.Logger) *Issuer {
	return &Issuer{
		keyRepository: kr,
		logger:        l,
	}
}

func (i Issuer) Create(ctx context.Context, req request.CreateIssuerRequest) (interface{}, error) {
	if req.DidMetadata.Method != "" && req.DidMetadata.Network != "" && req.DidMetadata.Blockchain != "" {
		//TODO set did issuer params accordingly
	}

	i.keyRepository.CreateKey(ctx)

	// Call sdk create issuer

	return "did", nil
}
