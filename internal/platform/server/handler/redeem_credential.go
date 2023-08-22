package handler

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/criteria"
	"bloock-identity-managed-api/internal/services/criteria/response"
	"errors"
	"github.com/gin-gonic/gin"
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

func RedeemCredential(credentialRedeem criteria.CredentialRedeem) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		proofs := ctx.QueryArray("proof")

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

		res, err := credentialRedeem.Redeem(ctx, bodyString, proofs)
		if err != nil {
			if errors.Is(domain.ErrInvalidProofType, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			if errors.Is(domain.ErrInvalidZkpMessage, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			if errors.Is(domain.ErrInvalidDID, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			if errors.Is(domain.ErrInvalidUUID, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			if errors.Is(domain.ErrInvalidCredentialSender, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, mapToRedeemCredentialResponse(res.(response.RedeemCredentialResponse)))
	}
}
