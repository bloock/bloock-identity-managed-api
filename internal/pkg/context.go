package pkg

import "context"

const (
	ApiKeyContextKey    = "API_KEY"
	IssuerKeyContextKey = "ISSUER_KEY"
)

func GetApiKeyFromContext(ctx context.Context) string {
	u, ok := ctx.Value(ApiKeyContextKey).(string)
	if !ok {
		return ""
	}
	return u
}

func GetIssuerKeyFromContext(ctx context.Context) string {
	u, ok := ctx.Value(IssuerKeyContextKey).(string)
	if !ok {
		return ""
	}
	return u
}
