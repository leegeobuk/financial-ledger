package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/financial-ledger/api"
	"github.com/leegeobuk/financial-ledger/cfg"
	"github.com/leegeobuk/financial-ledger/db"
	"github.com/spf13/viper"
)

func init() {
	profile := getProfile()
	log.Println("CONFIG_PROFILE:", profile)
	if err := loadConfig(profile); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	setGinMode(profile)
}

func getProfile() string {
	profile := os.Getenv("CONFIG_PROFILE")
	if len(profile) <= 0 {
		profile = "dev"
	}

	return profile
}

func loadConfig(profile string) error {
	viper.AddConfigPath("./cfg")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&cfg.Env)
}

func setGinMode(profile string) {
	switch profile {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "stg":
		gin.SetMode(gin.TestMode)
	case "prd":
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	idleConnsClosed, stopChan := make(chan struct{}), make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// create db
	dsn := cfg.Env.DB.DSN()
	mysql, err := db.NewMySQL(dsn)
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}

	const (
		interval = time.Second
		reps     = 30
	)
	if err = mysql.RetryPing(interval, reps); err != nil {
		log.Fatalf("Failed to ping db for %s: %v", reps*interval, err)
	}

	// create api server
	server := api.New(mysql)

	// gracefully shutdowns app when interrupt or terminal signal is received.
	go func() {
		select {
		case <-stopChan:
			log.Println("Got stop signal. Start graceful shutdown...")

			if err = mysql.Close(); err != nil {
				log.Printf("Error closing db: %v", err)
			}
			log.Println("DB closed")

			if err = server.Shutdown(); err != nil {
				log.Printf("Error while shutting down api server: %v", err)
			}
			log.Println("Server shutdown")

			log.Println("Gracefully shutdown. Bye.")
			close(idleConnsClosed)
		}
	}()

	if err = server.Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error running Server: %v", err)
	}

	<-idleConnsClosed
}
