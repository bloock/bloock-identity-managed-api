package issuer

import (
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/create/request"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type CreateIssuerRequest struct {
	Key             string      `json:"key" binding:"required"`
	DidMetadata     DidMetadata `json:"did_metadata"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Image           string      `json:"image"`
	PublishInterval int64       `json:"publish_interval"`
}

type DidMetadata struct {
	Method     string `json:"method"`
	Blockchain string `json:"blockchain"`
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

func CreateIssuer(l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateIssuerRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, api_error.NewBadRequestAPIError(err.Error()))
			return
		}

		createIssuerService := create.NewIssuer(ctx, req.Key, l)

		issuerReq := request.CreateIssuerRequest{
			Key: req.Key,
			DidMetadata: request.DidMetadataRequest{
				Method:     req.DidMetadata.Method,
				Blockchain: req.DidMetadata.Blockchain,
				Network:    req.DidMetadata.Network,
			},
			Name:            req.Name,
			Description:     req.Description,
			Image:           req.Image,
			PublishInterval: req.PublishInterval,
		}

		issuerDid, err := createIssuerService.Create(ctx, issuerReq)
		if err != nil {
			serverAPIError := api_error.NewInternalServerAPIError(err.Error())
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		ctx.JSON(http.StatusCreated, mapToCreateIssuerResponse(issuerDid))
	}
}
