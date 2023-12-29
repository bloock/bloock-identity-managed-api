package verification

import (
	"bloock-identity-managed-api/internal/domain"
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/platform/utils"
	"bloock-identity-managed-api/internal/services/criteria"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type CreateVerificationRequest struct {
	ZKRequest interface{} `json:"verification_json" binding:"required"`
	IssuerDid string      `json:"issuer_did" binding:"required"`
}

func CreateVerification(sym *utils.SyncMap, l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateVerificationRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, api_error.NewBadRequestAPIError(err.Error()))
			return
		}

		verificationService := criteria.NewCreateVerification(ctx, sym, l)

		request, err := verificationService.Create(ctx, req.ZKRequest, req.IssuerDid)
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
