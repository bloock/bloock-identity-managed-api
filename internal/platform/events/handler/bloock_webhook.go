package handler

import (
	"bloock-identity-managed-api/internal/platform/events/handler/action"
	"encoding/json"
	"errors"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func BloockWebhook(webhookSecret string, e action.ActionHandle) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		event, err := obtainEvent(ctx, webhookSecret)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if err = e.Dispatch(ctx, event); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}

func obtainEvent(ctx *gin.Context, secretKey string) (action.BloockEvent, error) {
	bloockSignature := ctx.GetHeader("Bloock-Signature")
	buf, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return action.BloockEvent{}, err
	}

	webhookClient := client.NewWebhookClient()
	ok, err := webhookClient.VerifyWebhookSignature(buf, bloockSignature, secretKey, false)
	if err != nil {
		return action.BloockEvent{}, err
	}
	if !ok {
		err = errors.New("invalid signature")
		return action.BloockEvent{}, err
	}
	var bloockEvent action.BloockEvent
	if err = json.Unmarshal(buf, &bloockEvent); err != nil {
		return action.BloockEvent{}, err
	}

	return bloockEvent, nil
}
