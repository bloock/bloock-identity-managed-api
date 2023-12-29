package verification

import (
	"bloock-identity-managed-api/internal/domain"
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/platform/utils"
	"bloock-identity-managed-api/internal/services/criteria"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
)

func CreateVerification(sym *utils.SyncMap, l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verificationJSON, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
			ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
			return
		}

		verificationService, err := criteria.NewCreateVerification(ctx, sym, l)
		if err != nil {
			badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
			ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
			return
		}

		request, err := verificationService.Create(ctx, verificationJSON)
		if err != nil {
			if errors.Is(domain.ErrInvalidVerificationRequest, err) {
				badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
				ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
				return
			}
			serverAPIError := api_error.NewInternalServerAPIError(err.Error())
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		ctx.Writer.Header().Set("Content-Type", "application/json")
		_, _ = ctx.Writer.Write(request)

	}
}
