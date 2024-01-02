package verification

import (
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/platform/utils"
	"bloock-identity-managed-api/internal/services/verify"
	"github.com/rs/zerolog"

	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type CallbackVerificationResponse struct {
	Success bool `json:"success"`
}

func CallbackVerification(sym *utils.SyncMap, l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId := ctx.Query("sessionId")
		if sessionId == "" {
			ctx.JSON(http.StatusBadRequest, "cannot proceed with an empty sessionId")
			return
		}

		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		verificationService, err := verify.NewVerificationCallback(ctx, sym, l)
		if err != nil {
			serverAPIError := api_error.NewInternalServerAPIError(err.Error())
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		err = verificationService.Verify(ctx, string(bodyBytes), sessionId)
		if err != nil {
			serverAPIError := api_error.NewInternalServerAPIError(err.Error())
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		ctx.JSON(http.StatusOK, CallbackVerificationResponse{Success: true})
	}
}
