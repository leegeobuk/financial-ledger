package testutil

import (
	"fmt"

	"github.com/leegeobuk/financial-ledger/cfg"
	"github.com/spf13/viper"
)

// SetupConfig is a utility function for setting config.
func SetupConfig(profile string) error {
	viper.AddConfigPath("../cfg")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	if err := viper.Unmarshal(&cfg.Env); err != nil {
		return fmt.Errorf("unmarshal envs to config: %w", err)
	}

	return nil
}
