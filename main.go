package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/financial-ledger/api"
	"github.com/leegeobuk/financial-ledger/cfg"
	"github.com/spf13/viper"
)

func init() {
	profile := getProfile()
	log.Println("CONFIG_PROFILE:", profile)
	switch profile {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "stg":
		gin.SetMode(gin.TestMode)
	case "prd":
		gin.SetMode(gin.ReleaseMode)
	}

	if err := initConfig(profile); err != nil {
		log.Fatalf("Error while loading config file: %v", err)
	}
}

func getProfile() string {
	profile := os.Getenv("CONFIG_PROFILE")
	if len(profile) <= 0 {
		profile = "dev"
	}

	return profile
}

func initConfig(profile string) error {
	viper.AddConfigPath("./cfg")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg.Env); err != nil {
		return err
	}

	return nil
}

func main() {
	ledgerAPI := api.New()
	ledgerAPI.Run()
}
