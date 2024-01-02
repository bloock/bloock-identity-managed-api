package pkg

import "context"

const (
	ApiKeyContextKey    = "API_KEY"
	EnvContextKey       = "ENV"
	IssuerDidContextKey = "ISSUER_DID"
	IssuerKeyContextKey = "ISSUER_KEY"
)

func GetApiKeyFromContext(ctx context.Context) string {
	u, ok := ctx.Value(ApiKeyContextKey).(string)
	if !ok {
		return ""
	}
	return u
}

func GetIssuerDidFromContext(ctx context.Context) string {
	u, ok := ctx.Value(IssuerDidContextKey).(string)
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

func GetEnvFromContext(ctx context.Context) *string {
	u, ok := ctx.Value(EnvContextKey).(string)
	if !ok {
		return nil
	}
	return &u
}
