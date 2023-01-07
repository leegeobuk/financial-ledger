package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server can start and stop api server.
type Server struct {
	port string
	host string

	server *http.Server
	router *gin.Engine
}

// New returns new Server struct.
func New(host, port string) *Server {
	router := gin.Default()

	return &Server{
		port: port,
		host: host,
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		router: router,
	}
}

// Run sets CORS and all handlers and then runs api server.
func (a *Server) Run() error {
	a.setCORS()
	a.setRoutes()

	return a.server.ListenAndServe()
}

func (a *Server) setCORS() gin.IRoutes {
	return a.router.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(a.host, ","),
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}

// Shutdown gracefully shutdowns api server
func (a *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	return a.server.Shutdown(ctx)
}
