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
		method := config.Configuration.Issuer.DidMetadata.Method
		blockchain := config.Configuration.Issuer.DidMetadata.Blockchain
		network := config.Configuration.Issuer.DidMetadata.Network

		issuerManagedKey, _ := c.GetQuery("issuer_key")
		if issuerManagedKey != "" {
			issuerKey = issuerManagedKey

			methodQuery, _ := c.GetQuery("method")
			blockchainQuery, _ := c.GetQuery("blockchain")
			networkQuery, _ := c.GetQuery("network")
			if method != "" {
				method = methodQuery
				blockchain = blockchainQuery
				network = networkQuery
			}
		}

		c.Set(pkg.IssuerDidTypeMethod, method)
		c.Set(pkg.IssuerDidTypeBlockchain, blockchain)
		c.Set(pkg.IssuerDidTypeNetwork, network)
		c.Set(pkg.IssuerKeyContextKey, issuerKey)

		c.Next()
	}
}
