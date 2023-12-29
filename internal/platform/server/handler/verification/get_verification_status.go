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

type GetVerificationStatusResponse struct {
	Success bool `json:"success"`
}

func GetVerificationStatus(sym *utils.SyncMap, l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId := ctx.Query("sessionId")
		if sessionId == "" {
			ctx.JSON(http.StatusBadRequest, "cannot proceed with an empty sessionId")
			return
		}

		verificationStatus := criteria.NewVerificationStatus(sym, l)

		err := verificationStatus.Get(ctx, sessionId)
		if err != nil {
			if errors.Is(domain.ErrSessionIdNotFound, err) {
				notFoundAPIError := api_error.NewAPIError(http.StatusNotFound, err.Error())
				ctx.JSON(notFoundAPIError.Status, notFoundAPIError)
				return
			}
			if errors.Is(domain.ErrNotVerified, err) {
				serverUnauthorized := api_error.NewUnauthorizedAPIError(err.Error())
				ctx.JSON(serverUnauthorized.Status, serverUnauthorized)
				return
			}
			if errors.Is(domain.ErrVerificationFailed, err) {
				serverUnauthorized := api_error.NewUnauthorizedAPIError(err.Error())
				ctx.JSON(serverUnauthorized.Status, serverUnauthorized)
				return
			}
			serverAPIError := api_error.NewInternalServerAPIError(err.Error())
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		ctx.JSON(http.StatusOK, GetVerificationStatusResponse{Success: true})
	}
}
