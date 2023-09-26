package handler

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/criteria"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetVerificationStatusResponse struct {
	Success bool `json:"success"`
}

func GetVerificationStatus(verification criteria.VerificationStatus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId := ctx.Query("sessionId")
		if sessionId == "" {
			ctx.JSON(http.StatusBadRequest, "cannot proceed with an empty sessionId")
			return
		}

		_, err := verification.Get(ctx, sessionId)
		if err != nil {
			if errors.Is(domain.ErrSessionIdNotFound, err) {
				ctx.JSON(http.StatusNotFound, NewBadRequestAPIError(err.Error()))
				return
			}
			if errors.Is(domain.ErrNotVerified, err) {
				ctx.JSON(http.StatusUnauthorized, NewBadRequestAPIError(err.Error()))
				return
			}
			if errors.Is(domain.ErrVerificationFailed, err) {
				ctx.JSON(http.StatusUnauthorized, NewBadRequestAPIError(err.Error()))
				return
			}
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, GetVerificationStatusResponse{Success: true})
	}
}
