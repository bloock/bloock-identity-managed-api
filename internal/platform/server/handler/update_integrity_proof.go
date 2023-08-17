package handler

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/update"
	"encoding/json"
	"errors"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebhookIntegrityProofRequest struct {
	WebhookId string                  `json:"webhook_id"`
	RequestId string                  `json:"request_id"`
	Type      string                  `json:"type"`
	CreatedAt int                     `json:"created_at"`
	Data      IntegrityProofConfirmed `json:"data"`
}

type IntegrityProofConfirmed struct {
	Leaves []string    `json:"leaves"`
	Nodes  []string    `json:"nodes"`
	Depth  string      `json:"depth"`
	Bitmap string      `json:"bitmap"`
	Anchor ProofAnchor `json:"anchor"`
	Type   string      `json:"type"`
}

type ProofAnchor struct {
	AnchorID int64           `json:"anchor_id"`
	Networks []AnchorNetwork `json:"networks"`
	Root     string          `json:"root"`
	Status   string          `json:"status"`
}

type AnchorNetwork struct {
	Name   string `json:"name"`
	State  string `json:"state"`
	TxHash string `json:"tx_hash"`
}

type UpdateIntegrityProofResponse struct {
	Success bool `json:"success"`
}

func UpdateIntegrityProof(updateIntegrityProof update.BloockIntegrityProofUpdate, secretKey string, enforceTolerance bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req WebhookIntegrityProofRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError("invalid json body"))
			return
		}
		bloockSignature := ctx.GetHeader("Bloock-Signature")

		webhookClient := client.NewWebhookClient()
		bodyBytes, err := json.Marshal(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}
		ok, err := webhookClient.VerifyWebhookSignature(bodyBytes, bloockSignature, secretKey, enforceTolerance)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}
		if !ok {
			ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError("invalid signature"))
			return
		}

		if err = updateIntegrityProof.Update(ctx, req.Data); err != nil {
			if errors.Is(domain.ErrInvalidIntegrityProof, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusAccepted, UpdateIntegrityProofResponse{Success: true})
	}
}
