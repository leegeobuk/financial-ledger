package cfg

import (
	"testing"

	"github.com/spf13/viper"
)

func TestDB_DSN(t *testing.T) {
	tests := []struct {
		name    string
		profile string
		want    string
	}{
		{
			name:    "success case: profile=local",
			profile: "local",
			want:    "user:1111@tcp(127.0.0.1:3306)/ledger",
		},
		{
			name:    "success case: profile=dev",
			profile: "dev",
			want:    "user:1111@tcp(ledger-db:3306)/ledger",
		},
		{
			name:    "success case: profile=stg",
			profile: "stg",
			want:    "user:1111@tcp(ledger-db:3306)/ledger",
		},
		{
			name:    "success case: profile=prd",
			profile: "prd",
			want:    "user:1111@tcp(ledger-db:3306)/ledger",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			if err := setupConfig(tt.profile); err != nil {
				t.Fatalf("Error setting up config: %v", err)
			}

			if got := Env.DB.DSN(); got != tt.want {
				t.Errorf("DSN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setupConfig(profile string) error {
	viper.AddConfigPath("../cfg")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&Env)
}
