package handler

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/criteria"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVerification(verification criteria.VerificationBySchemaId) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		schemaID := ctx.Param("id")
		if schemaID == "" {
			ctx.JSON(http.StatusBadRequest, "empty schema id")
			return
		}

		proofType := ctx.Query("proof")

		res, err := verification.Get(ctx, schemaID, proofType)
		if err != nil {
			if errors.Is(domain.ErrInvalidProofType, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.Writer.Header().Set("Content-Type", "application/json")
		_, _ = ctx.Writer.Write(res.([]byte))

	}
}
