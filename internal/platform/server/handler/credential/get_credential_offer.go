package credential

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/services/criteria"
	"bloock-identity-managed-api/internal/services/criteria/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type CredentialOfferResponse struct {
	ID       string                      `json:"id"`
	ThreadID string                      `json:"thid"`
	Body     CredentialOfferBodyResponse `json:"body"`
	From     string                      `json:"from"`
	To       string                      `json:"to"`
	Typ      string                      `json:"typ"`
	Type     string                      `json:"type"`
}

type CredentialOfferBodyResponse struct {
	URL         string            `json:"url"`
	Credentials []CredentialOffer `json:"credentials"`
}

type CredentialOffer struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func mapToCredentialOfferResponse(r response.GetCredentialOfferResponse) CredentialOfferResponse {
	credentialsOffers := make([]CredentialOffer, 0)
	credentialsOffers = append(credentialsOffers, CredentialOffer{
		ID:          r.Body.ID,
		Description: r.Body.Description,
	})
	body := CredentialOfferBodyResponse{
		URL:         r.Body.URL,
		Credentials: credentialsOffers,
	}
	res := CredentialOfferResponse{
		ID:       r.ID,
		ThreadID: r.ThreadID,
		Body:     body,
		From:     r.From,
		To:       r.To,
		Typ:      r.Typ,
		Type:     r.Type,
	}
	return res
}

func GetCredentialOffer(cr repository.CredentialRepository, l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		credentialId := ctx.Param("id")
		if credentialId == "" {
			ctx.JSON(http.StatusBadRequest, "empty credential id")
			return
		}

		credentialService, err := criteria.NewCredentialOffer(ctx, cr, l)

		offer, err := credentialService.Get(ctx, credentialId)
		if err != nil {
			if errors.Is(domain.ErrInvalidUUID, err) {
				badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
				ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
				return
			}
			serverAPIError := api_error.NewInternalServerAPIError(err.Error())
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		ctx.JSON(http.StatusOK, mapToCredentialOfferResponse(offer))
	}
}
