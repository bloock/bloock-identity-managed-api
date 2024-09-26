package server

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/domain/repository"
	handler2 "bloock-identity-managed-api/internal/platform/events/handler"
	"bloock-identity-managed-api/internal/platform/events/handler/action"
	"bloock-identity-managed-api/internal/platform/server/handler"
	"bloock-identity-managed-api/internal/platform/server/handler/credential"
	"bloock-identity-managed-api/internal/platform/server/handler/issuer"
	"bloock-identity-managed-api/internal/platform/server/handler/verification"
	"bloock-identity-managed-api/internal/platform/server/middleware"
	"bloock-identity-managed-api/internal/platform/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	host   string
	port   string
	engine *gin.Engine
	logger zerolog.Logger
}

func NewServer(cr repository.CredentialRepository, vm, au *utils.SyncMap, webhookSecretKey string, l zerolog.Logger) (*Server, error) {
	l = l.With().Str("layer", "infrastructure").Str("component", "gin").Logger()
	gin.DefaultWriter = l.With().Str("level", "info").Logger()
	gin.DefaultErrorWriter = l.With().Str("level", "error").Logger()

	router := gin.Default()
	if config.Configuration.Api.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	if err := router.SetTrustedProxies(nil); err != nil {
		return nil, err
	}
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))
	router.Use(middleware.ErrorMiddleware())
	router.Use(logger.SetLogger(
		logger.WithSkipPath([]string{"/health"}),
		logger.WithUTC(true),
		logger.WithLogger(func(c *gin.Context, _ zerolog.Logger) zerolog.Logger {
			return l
		}),
	))

	v1 := router.Group("/v1")
	v1.GET("health", handler.Health())

	v1.POST("/issuers", middleware.AuthMiddleware(), issuer.CreateIssuer(l))
	v1.GET("/issuers", issuer.GetIssuer())
	v1.POST("/issuers/state/publish", middleware.AuthMiddleware(), middleware.IssuerMiddleware(l), issuer.PublishIssuerState(l))

	v1.POST("/credentials", middleware.AuthMiddleware(), middleware.IssuerMiddleware(l), credential.CreateCredential(cr, l))
	v1.POST("/credentials/redeem", credential.RedeemCredential(cr, au, l))
	v1.GET("/credentials/:id/offer", middleware.AuthMiddleware(), middleware.IssuerMiddleware(l), credential.GetCredentialOffer(cr, au, l))
	v1.GET("/credentials/:id", credential.GetCredentialById(cr, l))
	v1.PUT("/credentials/:id/revocation", middleware.AuthMiddleware(), middleware.IssuerMiddleware(l), credential.RevokeCredential(cr, l))

	v1.POST("/verifications", middleware.AuthMiddleware(), verification.CreateVerification(vm, au, l))
	v1.POST("/verifications/callback", verification.CallbackVerification(vm, au, l))
	v1.GET("/verifications/status", verification.GetVerificationStatus(vm, l))

	actionHandler := action.NewActionHandle()
	sparseMtProof := action.NewSparseMtProofConfirmed(cr, l)

	actionHandler.Register(sparseMtProof.EventType(), sparseMtProof)

	router.POST("/bloock-events", handler2.BloockWebhook(webhookSecretKey, actionHandler))

	return &Server{
		host:   config.Configuration.Api.Host,
		port:   config.Configuration.Api.Port,
		engine: router,
		logger: l,
	}, nil
}

func (s *Server) Start() error {
	if err := s.engine.Run(fmt.Sprintf("%s:%s", s.host, s.port)); err != nil {
		return err
	}
	return nil
}
