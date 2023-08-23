package handler

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/create/request"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCredentialRequest struct {
	SchemaId          string              `json:"schema_id" binding:"required"`
	SchemaType        string              `json:"schema_type" binding:"required"`
	HolderDid         string              `json:"holder_did" binding:"required"`
	CredentialSubject []CredentialSubject `json:"credential_subject" binding:"required"`
	Expiration        int64               `json:"expiration" binding:"required"`
	Version           int32               `json:"version"`
}

type CredentialSubject struct {
	DataType string      `json:"data_type" binding:"required"`
	Key      string      `json:"key" binding:"required"`
	Value    interface{} `json:"value" binding:"required"`
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

		var req CreateCredentialRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
			return
		}

		credentialSubject := make([]request.CredentialSubject, 0)
		for _, cs := range req.CredentialSubject {
			credentialSubject = append(credentialSubject, request.CredentialSubject{
				DataType: cs.DataType,
				Key:      cs.Key,
				Value:    cs.Value,
			})
		}
		cr := request.CredentialRequest{
			SchemaId:          req.SchemaId,
			SchemaType:        req.SchemaType,
			HolderDid:         req.HolderDid,
			CredentialSubject: credentialSubject,
			Expiration:        req.Expiration,
			Version:           req.Version,
			Proofs:            proofs,
		}
		res, err := credential.Create(ctx, cr)
		if err != nil {
			if errors.Is(domain.ErrInvalidDataType, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, mapToCreateCredentialResponse(res.(string)))
	}
}
