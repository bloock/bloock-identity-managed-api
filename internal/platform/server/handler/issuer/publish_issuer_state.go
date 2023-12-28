package issuer

import (
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/services/publish"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type PublishIssuerStateResponse struct {
	TxHash string `json:"tx_hash"`
}

func mapToPublishIssuerStateResponse(hash string) PublishIssuerStateResponse {
	return PublishIssuerStateResponse{
		TxHash: hash,
	}
}

func PublishIssuerState(l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		issuerService, err := publish.NewIssuerPublish(ctx, l)
		if err != nil {
			badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
			ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
			return
		}

		txHash, err := issuerService.Publish(ctx)
		if err != nil {
			serverAPIError := api_error.NewInternalServerAPIError(err.Error())
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		ctx.JSON(http.StatusOK, mapToPublishIssuerStateResponse(txHash))
	}
}
