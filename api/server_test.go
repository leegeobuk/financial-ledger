package api

import (
	"log"
	"net"
	"reflect"
	"testing"

	"github.com/leegeobuk/financial-ledger/cfg"
	"github.com/spf13/viper"
)

func TestAPI_Run(t *testing.T) {
	// given
	if err := setup("dev"); err != nil {
		log.Fatalf("Error setting up: %v", err)
	}
	host := cfg.Env.Server.Host

	tests := []struct {
		name    string
		server  *Server
		wantErr error
	}{
		{
			name:    "fail case: port=xxx",
			server:  New(host, "xxx"),
			wantErr: &net.OpError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.server.Run()
			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("Run() error = %T, wantErr %T", err, tt.wantErr)
			}
		})
	}
}

func TestServer_Shutdown(t *testing.T) {
	// given
	if err := setup("dev"); err != nil {
		log.Fatalf("Error setting up: %v", err)
	}
	host := cfg.Env.Server.Host
	port := cfg.Env.Server.Port

	tests := []struct {
		name    string
		server  *Server
		wantErr error
	}{
		{
			name:    "success case: no error returned",
			server:  New(host, port),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				tt.server.Run()
			}()

			err := tt.server.Shutdown()
			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func setup(profile string) error {
	viper.AddConfigPath("../cfg")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&cfg.Env)
}
