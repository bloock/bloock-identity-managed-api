package server

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/gin-gonic/gin"
)

type Server struct {
	host   string
	port   string
	engine *gin.Engine
	debug  bool
	logger zerolog.Logger
}

func NewServer(host string, port string, webhookSecretKey string, enforceTolerance bool, logger zerolog.Logger, debug bool) (*Server, error) {
	router := gin.Default()
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	if err := router.SetTrustedProxies(nil); err != nil {
		return nil, err
	}

	//v1 := router.Group("/v1/")

	return &Server{
		host: host,
		port: port,
		engine: router,
		debug: debug,
		logger: logger,
	}, nil
}

func (s *Server) Start() error {
	if err := s.engine.Run(fmt.Sprintf("%s:%s", s.host, s.port)); err != nil {
		return err
	}
	return nil
}
