package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupServer(host, port string) (testServer *Server) {
	router := gin.Default()

	testServer = &Server{
		port: port,
		host: host,
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		router: router,
	}

	return
}
