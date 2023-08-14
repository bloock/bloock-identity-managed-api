package handler

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/criteria"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCredentialById(credential criteria.CredentialById) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		issuerDid := ctx.Param("issuer_did")
		if issuerDid == "" {
			ctx.JSON(http.StatusBadRequest, "empty issuer did")
			return
		}

		credentialId := ctx.Param("credential_id")
		if credentialId == "" {
			ctx.JSON(http.StatusBadRequest, "empty credential id")
			return
		}

		res, err := credential.Get(ctx, issuerDid, credentialId)
		if err != nil {
			if errors.Is(domain.ErrInvalidUUID, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			if errors.Is(domain.ErrInvalidDID, err) {
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
		cred, ok := res.(map[string]interface{})
		if !ok {
			ctx.JSON(http.StatusNotFound, domain.ErrCredentialNotFound.Error())
			return
		}

		ctx.JSON(http.StatusOK, cred)
	}
}
