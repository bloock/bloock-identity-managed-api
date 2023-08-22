package handler

import (
	"bloock-identity-managed-api/internal/services/criteria"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetIssuerListResponse struct {
	Issuers []string `json:"issuers"`
}

func mapToGetIssuerListResponse(issuers []string) GetIssuerListResponse {
	return GetIssuerListResponse{
		Issuers: issuers,
	}
}

func GetIssuerList(issuer criteria.IssuerList) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := issuer.Get(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, mapToGetIssuerListResponse(res.([]string)))
	}
}
