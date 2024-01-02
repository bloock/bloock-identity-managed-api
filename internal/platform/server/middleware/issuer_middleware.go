package middleware

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/pkg"
	api_error "bloock-identity-managed-api/internal/platform/server/error"
	"bloock-identity-managed-api/internal/services/create/request"
	"bloock-identity-managed-api/internal/services/criteria"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func IssuerMiddleware(l zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		issuerKey, _ := c.GetQuery("issuer_key")

		if issuerKey != "" {
			method, _ := c.GetQuery("method")
			blockchain, _ := c.GetQuery("blockchain")
			network, _ := c.GetQuery("network")

			cm := config.Configuration.Issuer.DidMetadata
			mr := request.DidMetadataRequest{Method: cm.Method, Blockchain: cm.Blockchain, Network: cm.Network}
			if method != "" {
				mr.Method = method
				mr.Blockchain = blockchain
				mr.Network = network
			}
			getIssuerService := criteria.NewIssuerByKey(c, issuerKey, l)
			issuerDid, err := getIssuerService.Get(c, mr)
			if err != nil {
				serverAPIError := api_error.NewInternalServerAPIError(err.Error())
				c.JSON(serverAPIError.Status, serverAPIError)
				c.Abort()
			}
			c.Set(pkg.IssuerDidContextKey, issuerDid)
			c.Set(pkg.IssuerKeyContextKey, issuerKey)
		} else {
			c.Set(pkg.IssuerDidContextKey, config.Configuration.Issuer.IssuerDid)
			c.Set(pkg.IssuerKeyContextKey, config.Configuration.Issuer.Key.Key)
		}

		c.Next()
	}
}
