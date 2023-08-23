package handler

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/criteria"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iden3/go-schema-processor/verifiable"
	"net/http"
)

func GetCredentialById(credential criteria.CredentialById) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		credentialId := ctx.Param("credential_id")
		if credentialId == "" {
			ctx.JSON(http.StatusBadRequest, "empty credential id")
			return
		}

		res, err := credential.Get(ctx, credentialId)
		if err != nil {
			if errors.Is(domain.ErrInvalidUUID, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			if errors.Is(domain.ErrCredentialNotFound, err) {
				ctx.JSON(http.StatusNotFound, domain.ErrCredentialNotFound.Error())
				return
			}
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}
		cred, ok := res.(verifiable.W3CCredential)
		if !ok {
			ctx.JSON(http.StatusNotFound, domain.ErrCredentialNotFound.Error())
			return
		}

		ctx.JSON(http.StatusOK, cred)
	}
}
