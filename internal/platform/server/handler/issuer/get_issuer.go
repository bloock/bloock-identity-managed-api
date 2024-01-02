package issuer

import (
	"bloock-identity-managed-api/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetIssuerResponse struct {
	Did string `json:"did"`
}

func mapToGetIssuerResponse(did string) GetIssuerResponse {
	return GetIssuerResponse{
		Did: did,
	}
}

func GetIssuer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, mapToGetIssuerResponse(config.Configuration.Issuer.IssuerDid))
	}
}
