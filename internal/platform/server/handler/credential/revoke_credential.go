package credential

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/services/cancel"
	"bloock-identity-managed-api/internal/services/criteria"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type RevokeCredentialResponse struct {
	Success bool `json:"success"`
}

func RevokeCredential(cr repository.CredentialRepository, l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		credentialId := ctx.Param("id")
		if credentialId == "" {
			ctx.JSON(http.StatusBadRequest, "empty credential id")
			return
		}

		credentialService := criteria.NewCredentialById(cr, l)
		revokeService, err := cancel.NewCredentialRevocation(ctx, l)
		if err != nil {
			badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
			ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
			return
		}

		cred, err := credentialService.Get(ctx, credentialId)
		if err != nil {
			if errors.Is(domain.ErrInvalidUUID, err) {
				badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
				ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
				return
			}
			if errors.Is(domain.ErrCredentialNotFound, err) {
				notFoundAPIError := api_error.NewAPIError(http.StatusNotFound, err.Error())
				ctx.JSON(notFoundAPIError.Status, notFoundAPIError)
				return
			}
			serverAPIError := api_error.NewInternalServerAPIError(err.Error())
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		if err = revokeService.Revoke(ctx, cred); err != nil {
			serverAPIError := api_error.NewInternalServerAPIError(err.Error())
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		ctx.JSON(http.StatusOK, RevokeCredentialResponse{Success: true})
	}
}
