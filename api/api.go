package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/financial-ledger/cfg"
)

// API can start and stop api server.
type API struct {
	port string
	host string

	server *http.Server
	router *gin.Engine
}

// New returns new API struct.
func New() *API {
	port := cfg.Env.Server.Port
	router := gin.Default()

	return &API{
		port: port,
		host: cfg.Env.Server.Host,
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		router: router,
	}
}

// Run sets CORS and all handlers and then runs api server.
func (a *API) Run() {
	a.router.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(a.host, ","),
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	a.setRoutes()

	if err := a.server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("run api: %v", err)
	}
}

// Shutdown gracefully shutdowns api server
func (a *API) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := a.server.Shutdown(ctx); errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("api already shutdown: %v", err)
	}
}
