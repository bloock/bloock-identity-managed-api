package handler

import (
	"bloock-identity-managed-api/internal/services/criteria"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetIssuerResponse struct {
	IssuerDid string `json:"issuer_did"`
}

func mapToGetIssuerResponse(did string) GetIssuerResponse {
	return GetIssuerResponse{
		IssuerDid: did,
	}
}

func GetIssuer(issuer criteria.Issuer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := issuer.Get(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, mapToGetIssuerResponse(res.(string)))
	}
}
