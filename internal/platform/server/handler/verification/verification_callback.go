package verification

import (
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/services/verify"

	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type VerificationCallbackResponse struct {
	Success bool `json:"success"`
}

func VerificationCallback(verification verify.VerificationCallback) gin.HandlerFunc {
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

		_, err = verification.Verify(ctx, string(bodyBytes), sessionId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, api_error.NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, VerificationCallbackResponse{Success: true})
	}
}
