package credential

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/platform/utils"
	"bloock-identity-managed-api/internal/services/criteria"
	"bloock-identity-managed-api/internal/services/criteria/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

type RedeemCredentialResponse struct {
	Body     IssuanceMessageBody `json:"body"`
	From     string              `json:"from"`
	Id       string              `json:"id"`
	ThreadID string              `json:"threadID"`
	To       string              `json:"to"`
	Typ      string              `json:"typ"`
	Type     string              `json:"type"`
}

type IssuanceMessageBody struct {
	Credential interface{} `json:"credential"`
}

func mapToRedeemCredentialResponse(res response.RedeemCredentialResponse) RedeemCredentialResponse {
	return RedeemCredentialResponse{
		Body: IssuanceMessageBody{
			Credential: res.Body,
		},
		From:     res.From,
		To:       res.To,
		Id:       res.ID,
		ThreadID: res.ThreadID,
		Typ:      res.Typ,
		Type:     res.Type,
	}
}

func RedeemCredential(cr repository.CredentialRepository, au *utils.SyncMap, l zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		threadID := ctx.Query("thread_id")
		if threadID == "" {
			ctx.JSON(http.StatusBadRequest, "cannot proceed with an empty thread id")
			return
		}

		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		bodyString := string(bodyBytes)
		if bodyString == "" {
			ctx.JSON(http.StatusBadRequest, "cannot proceed with an empty request")
			return
		}

		credentialService, err := criteria.NewCredentialRedeem(ctx, cr, au, threadID, l)
		if err != nil {
			badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
			ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
			return
		}

		res, err := credentialService.Redeem(ctx, bodyString)
		if err != nil {
			if errors.Is(domain.ErrInvalidZkpMessage, err) {
				badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
				ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
				return
			}
			if errors.Is(domain.ErrInvalidDID, err) {
				badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
				ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
				return
			}
			if errors.Is(domain.ErrInvalidUUID, err) {
				badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
				ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
				return
			}
			if errors.Is(domain.ErrInvalidCredentialSender, err) {
				badRequestAPIError := api_error.NewBadRequestAPIError(err.Error())
				ctx.JSON(badRequestAPIError.Status, badRequestAPIError)
				return
			}
			serverAPIError := api_error.NewInternalServerAPIError(err)
			ctx.JSON(serverAPIError.Status, serverAPIError)
			return
		}

		ctx.JSON(http.StatusOK, mapToRedeemCredentialResponse(res))
	}
}
