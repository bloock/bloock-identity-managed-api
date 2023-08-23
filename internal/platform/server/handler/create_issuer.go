package handler

import (
	"bloock-identity-managed-api/internal/services/create"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateIssuerRequest struct {
	DidMetadataRequest `json:"did_metadata"`
}

type DidMetadataRequest struct {
	Blockchain string `json:"blockchain"`
	Method     string `json:"method"`
	Network    string `json:"network"`
}

type CreateIssuerResponse struct {
	Did string `json:"did"`
}

func mapToCreateIssuerResponse(did string) CreateIssuerResponse {
	return CreateIssuerResponse{
		Did: did,
	}
}

func CreateIssuer(issuer create.Issuer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateIssuerRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
			return
		}

		/*ir := request.CreateIssuerRequest{
			DidMetadata: request.DidMetadataRequest{
				Method:     req.DidMetadataRequest.Method,
				Blockchain: req.DidMetadataRequest.Blockchain,
				Network:    req.DidMetadataRequest.Network,
			},
		}*/

		res, err := issuer.Create(ctx, req.DidMetadataRequest.Method, req.DidMetadataRequest.Blockchain, req.DidMetadataRequest.Network)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, mapToCreateIssuerResponse(res.(string)))
	}
}
