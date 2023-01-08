package testutil

import (
	"github.com/leegeobuk/financial-ledger/cfg"
	"github.com/spf13/viper"
)

// SetupConfig is a utility function for setting config.
func SetupConfig(profile string) error {
	viper.AddConfigPath("../cfg")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&cfg.Env)
}
