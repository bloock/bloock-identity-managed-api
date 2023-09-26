package server

import (
	handler2 "bloock-identity-managed-api/internal/platform/events/handler"
	"bloock-identity-managed-api/internal/platform/events/handler/action"
	"bloock-identity-managed-api/internal/platform/server/handler"
	"bloock-identity-managed-api/internal/services/cancel"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/criteria"
	"bloock-identity-managed-api/internal/services/publish"
	"bloock-identity-managed-api/internal/services/update"
	"bloock-identity-managed-api/internal/services/verify"
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

func NewServer(host string, port string, c create.Credential, co criteria.CredentialOffer, rc criteria.CredentialRedeem, cbi criteria.CredentialById, bpu update.IntegrityProofUpdate, smp update.SparseMtProofUpdate,
	gi criteria.Issuer, crv cancel.CredentialRevocation, pi publish.IssuerPublish, vbs criteria.VerificationBySchemaId, vc verify.VerificationCallback, vs criteria.VerificationStatus,
	webhookSecretKey string, logger zerolog.Logger, debug bool) (*Server, error) {
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

	v1.GET("/issuers", handler.GetIssuer(gi))
	v1.POST("/issuers/state/publish", handler.PublishIssuerState(pi))

	v1.POST("/claims", handler.CreateCredential(c))
	v1.POST("/claims/redeem", handler.RedeemCredential(rc))

	v1.GET("/claims/:credential_id/offer", handler.GetCredentialOffer(co))
	v1.GET("/claims/:credential_id", handler.GetCredentialById(cbi))
	v1.PUT("/claims/:credential_id/revoke", handler.RevokeCredential(cbi, crv))

	v1.GET("/claims/:schema_id/verification", handler.GetVerification(vbs))
	v1.POST("/verification/callback", handler.VerificationCallback(vc))
	v1.GET("/verification/status", handler.GetVerificationStatus(vs))

	actionHandler := action.NewActionHandle()

	integrityProof := action.NewIntegrityProofConfirmed(bpu, logger)
	sparseMtProof := action.NewSparseMtProofConfirmed(smp, logger)

	actionHandler.Register(integrityProof.EventType(), integrityProof)
	actionHandler.Register(sparseMtProof.EventType(), sparseMtProof)

	router.POST("/bloock-events", handler2.BloockWebhook(webhookSecretKey, actionHandler))

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
