package middleware

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func IssuerMiddleware(l zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		issuerKey := config.Configuration.Issuer.Key.Key

		issuerManagedKey, _ := c.GetQuery("issuer_key")
		if issuerManagedKey != "" {
			issuerKey = issuerManagedKey
		}

		c.Set(pkg.IssuerKeyContextKey, issuerKey)

		c.Next()
	}
}
