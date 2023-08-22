package server

import (
	"bloock-identity-managed-api/internal/platform/server/handler"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/criteria"
	"bloock-identity-managed-api/internal/services/update"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	host   string
	port   string
	engine *gin.Engine
	debug  bool
	logger zerolog.Logger
}

func NewServer(host string, port string, c create.Credential, co criteria.CredentialOffer, rc criteria.CredentialRedeem, ci criteria.CredentialById, bpu update.BloockIntegrityProofUpdate,
	webhookSecretKey string, enforceTolerance bool, logger zerolog.Logger, debug bool) (*Server, error) {
	router := gin.Default()
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	if err := router.SetTrustedProxies(nil); err != nil {
		return nil, err
	}

	v1 := router.Group("/v1")

	v1.POST("/:issuer_did/credentials", handler.CreateCredential(c))
	v1.POST("/credentials/redeem", handler.RedeemCredential(rc))

	v1.GET("/credentials/:credential_id/offer", handler.GetCredentialOffer(co))
	v1.GET("/:issuer_did/credentials/:credential_id", handler.GetCredentialById(ci))

	v1.POST("/integrity/webhook", handler.UpdateIntegrityProof(bpu, webhookSecretKey, enforceTolerance))

	return &Server{
		host:   host,
		port:   port,
		engine: router,
		debug:  debug,
		logger: logger,
	}, nil
}

func (s *Server) Start() error {
	if err := s.engine.Run(fmt.Sprintf("%s:%s", s.host, s.port)); err != nil {
		return err
	}
	return nil
}
