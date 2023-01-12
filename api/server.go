package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/cfg"
	"github.com/leegeobuk/household-ledger/db"
)

// Server can start and stop api server.
type Server struct {
	port string
	host string

	server *http.Server
	router *gin.Engine

	db *db.MySQL
}

// New returns new Server struct.
func New(mysql *db.MySQL) *Server {
	host, port := cfg.Env.Server.Host, cfg.Env.Server.Port
	router := gin.Default()

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: time.Second * 10,
	}

	return &Server{
		port:   port,
		host:   host,
		server: server,
		router: router,
		db:     mysql,
	}
}

// Run sets CORS and all handlers and then runs api server.
func (s *Server) Run() error {
	s.setCORS()
	s.setRoutes()

	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("run api server: %w", err)
	}

	return nil
}

func (s *Server) setCORS() gin.IRoutes {
	return s.router.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(s.host, ","),
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}

// Shutdown gracefully shutdowns api server.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown api server: %w", err)
	}

	return nil
}
