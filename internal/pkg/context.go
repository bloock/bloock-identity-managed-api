package pkg

import "context"

const (
	ApiKeyContextKey        = "API_KEY"
	EnvContextKey           = "ENV"
	IssuerDidTypeMethod     = "ISSUER_DID_TYPE_METHOD"
	IssuerDidTypeBlockchain = "ISSUER_DID_TYPE_BLOCKCHAIN"
	IssuerDidTypeNetwork    = "ISSUER_DID_TYPE_NETWORK"
	IssuerKeyContextKey     = "ISSUER_KEY"
)

func GetApiKeyFromContext(ctx context.Context) string {
	u, ok := ctx.Value(ApiKeyContextKey).(string)
	if !ok {
		return ""
	}
	return u
}

func GetIssuerDidTypeMethodFromContext(ctx context.Context) string {
	u, ok := ctx.Value(IssuerDidTypeMethod).(string)
	if !ok {
		return ""
	}
	return u
}

func GetIssuerDidTypeBlockchainFromContext(ctx context.Context) string {
	u, ok := ctx.Value(IssuerDidTypeBlockchain).(string)
	if !ok {
		return ""
	}
	return u
}

func GetIssuerDidTypeNetworkFromContext(ctx context.Context) string {
	u, ok := ctx.Value(IssuerDidTypeNetwork).(string)
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
