package handler

import (
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/create/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCredentialRequest struct {
	SchemaId          string                 `json:"schema_id" binding:"required"`
	SchemaType        string                 `json:"schema_type" binding:"required"`
	HolderDid         string                 `json:"holder_did" binding:"required"`
	CredentialSubject map[string]interface{} `json:"credential_subject" binding:"required"`
	Expiration        int64                  `json:"expiration" binding:"required"`
	Version           int32                  `json:"version"`
}

type CreateCredentialResponse struct {
	Id string `json:"id"`
}

func mapToCreateCredentialResponse(id string) CreateCredentialResponse {
	return CreateCredentialResponse{
		Id: id,
	}
}

func CreateCredential(credential create.Credential) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		proofs := ctx.QueryArray("proof")

		issuerDid := ctx.Param("issuer_did")
		if issuerDid == "" {
			ctx.JSON(http.StatusBadRequest, "empty issuer did")
			return
		}

		var req CreateCredentialRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
			return
		}

		cr := request.CredentialRequest{
			SchemaId:          req.SchemaId,
			SchemaType:        req.SchemaType,
			IssuerDid:         issuerDid,
			HolderDid:         req.HolderDid,
			CredentialSubject: req.CredentialSubject,
			Expiration:        req.Expiration,
			Version:           req.Version,
			Proofs:            proofs,
		}
		res, err := credential.Create(ctx, cr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, mapToCreateCredentialResponse(res.(string)))
	}
}
