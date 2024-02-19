package create

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/platform/identity"
	keyRepo "bloock-identity-managed-api/internal/platform/key"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"github.com/rs/zerolog"
	"log"
)

type Issuer struct {
	keyRepository      repository.KeyRepository
	identityRepository repository.IdentityRepository
	logger             zerolog.Logger
}

func NewIssuer(ctx context.Context, key string, l zerolog.Logger) *Issuer {
	return &Issuer{
		keyRepository:      keyRepo.NewKeyRepository(ctx, key, l),
		identityRepository: identity.NewIdentityRepository(ctx, l),
		logger:             l,
	}
}

func (i Issuer) Create(ctx context.Context, req request.CreateIssuerRequest) (string, error) {
	didType, err := domain.GetDidType(req.DidMetadata.Method, req.DidMetadata.Blockchain, req.DidMetadata.Network)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return "", err
	}

	issuerKey, err := i.keyRepository.LoadIssuerKey(ctx)
	if err != nil {
		log.Println("Enter Import issuer")
		return "", err
	}

	issuer, err := i.identityRepository.ImportIssuer(ctx, issuerKey, didType)
	if err != nil {
		log.Println("Enter Import issuer")
		return "", err
	}
	if issuer.Did.Did != "" {
		return issuer.Did.Did, nil
	}

	publishInterval, err := domain.NewPublishIntervalMinutes(req.PublishInterval)
	if err != nil {
		i.logger.Error().Err(err).Msg("")
		return "", err
	}

	return i.identityRepository.CreateIssuer(ctx, issuerKey, didType, req.Name, req.Description, req.Image, publishInterval)
}
