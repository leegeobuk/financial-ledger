package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/financial-ledger/cfg"
	"github.com/spf13/viper"
)

func setupConfig(profile string) error {
	viper.AddConfigPath("../cfg")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&cfg.Env)
}

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
