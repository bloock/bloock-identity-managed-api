package credential

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/services/criteria"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

func GetCredentialById(cr repository.CredentialRepository, l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		credentialId := ctx.Param("id")
		if credentialId == "" {
			ctx.JSON(http.StatusBadRequest, "empty credential id")
			return
		}

		credentialService := criteria.NewCredentialById(cr, l)

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

		ctx.JSON(http.StatusOK, cred)
	}
}
