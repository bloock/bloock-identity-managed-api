package issuer

import (
	"bloock-identity-managed-api/internal/domain"
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/create/request"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"mime/multipart"
	"net/http"
)

type CreateIssuerRequest struct {
	Key             string                `form:"key" binding:"required"`
	Name            string                `form:"name"`
	Description     string                `form:"description"`
	Image           *multipart.FileHeader `form:"image"`
	PublishInterval int                   `form:"publish_interval"`
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
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, api_error.NewBadRequestAPIError(err.Error()))
			return
		}

		var image []byte
		var err error
		if req.Image != nil {
			image, err = loadImage(req.Image)
			if err != nil {
				badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
				ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
				return
			}
		}

		createIssuerService := create.NewIssuer(ctx, req.Key, l)

		issuerReq := request.CreateIssuerRequest{
			Key:             req.Key,
			Name:            req.Name,
			Description:     req.Description,
			Image:           base64.URLEncoding.EncodeToString(image),
			PublishInterval: req.PublishInterval,
		}

		issuerDid, err := createIssuerService.Create(ctx, issuerReq)
		if err != nil {
			if errors.Is(domain.ErrInvalidPublishIntervalMinutes, err) {
				badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
				ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
				return
			}
			serverAPIError := api_error.NewInternalServerAPIError(err)
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		ctx.JSON(http.StatusCreated, mapToCreateIssuerResponse(issuerDid))
	}
}

func loadImage(image *multipart.FileHeader) ([]byte, error) {
	imageReader, err := image.Open()
	if err != nil {
		return nil, err
	}

	file, err := io.ReadAll(imageReader)
	if err != nil {
		return nil, err
	}
	if len(file) == 0 {
		return nil, fmt.Errorf("file must be a valid file")
	}

	return file, nil
}
