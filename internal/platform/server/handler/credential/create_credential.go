package credential

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/create/request"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type CreateCredentialRequest struct {
	SchemaId          string              `json:"schema_id" binding:"required"`
	HolderDid         string              `json:"holder_did" binding:"required"`
	CredentialSubject []CredentialSubject `json:"credential_subject" binding:"required"`
	Expiration        int64               `json:"expiration" binding:"required"`
	Version           int32               `json:"version"`
}

type CredentialSubject struct {
	Key   string      `json:"key" binding:"required"`
	Value interface{} `json:"value" binding:"required"`
}

type CreateCredentialResponse struct {
	Id string `json:"id"`
}

func mapToCreateCredentialResponse(id string) CreateCredentialResponse {
	return CreateCredentialResponse{
		Id: id,
	}
}

func CreateCredential(cr repository.CredentialRepository, l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateCredentialRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, api_error.NewBadRequestAPIError(err.Error()))
			return
		}

		credentialService, err := create.NewCredential(ctx, cr, l)
		if err != nil {
			badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
			ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
			return
		}

		credentialSubject := make([]request.CredentialSubject, 0)
		for _, cs := range req.CredentialSubject {
			credentialSubject = append(credentialSubject, request.CredentialSubject{
				Key:   cs.Key,
				Value: cs.Value,
			})
		}
		cdr := request.CredentialRequest{
			SchemaId:          req.SchemaId,
			HolderDid:         req.HolderDid,
			CredentialSubject: credentialSubject,
			Expiration:        req.Expiration,
			Version:           req.Version,
		}
		credentialID, err := credentialService.Create(ctx, cdr)
		if err != nil {
			if errors.Is(domain.ErrInvalidDataType, err) {
				badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
				ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
				return
			}
			serverAPIError := api_error.NewInternalServerAPIError(err)
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		ctx.JSON(http.StatusCreated, mapToCreateCredentialResponse(credentialID))
	}
}
