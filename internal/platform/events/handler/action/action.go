package action

import "context"

type BloockEvent struct {
	WebhookId string      `json:"webhook_id"`
	RequestId string      `json:"request_id"`
	Type      string      `json:"type"`
	CreatedAt int         `json:"created_at"`
	Data      interface{} `json:"data"`
}

type EventAction interface {
	EventType() string
	Run(ctx context.Context, bloockEvent BloockEvent) error
}

type ActionHandle struct {
	handlers map[string]EventAction
}

func NewActionHandle() ActionHandle {
	return ActionHandle{handlers: make(map[string]EventAction)}
}

func (a ActionHandle) Register(t string, e EventAction) {
	a.handlers[t] = e
}

func (a ActionHandle) Dispatch(ctx context.Context, bloockEvent BloockEvent) error {
	handler, ok := a.handlers[bloockEvent.Type]
	if !ok {
		return nil
	}
	return handler.Run(ctx, bloockEvent)
}
