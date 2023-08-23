package handler

import (
	"bloock-identity-managed-api/internal/services/publish"
	"github.com/gin-gonic/gin"
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

func PublishIssuerState(issuer publish.IssuerPublish) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := issuer.Publish(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, mapToPublishIssuerStateResponse(res.(string)))
	}
}
