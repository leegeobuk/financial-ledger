package api

import (
	"log"
	"net"
	"reflect"
	"testing"

	"github.com/leegeobuk/financial-ledger/cfg"
)

func TestServer_Run(t *testing.T) {
	// given
	if err := setupConfig("dev"); err != nil {
		log.Fatalf("Error setting up: %v", err)
	}
	host := cfg.Env.Server.Host

	tests := []struct {
		name       string
		server     *Server
		shouldFail bool
		wantErr    error
	}{
		{
			name:       "fail case: port=xxx",
			server:     setupServer(host, "xxx"),
			shouldFail: true,
			wantErr:    &net.OpError{},
		},
		{
			name:       "success case: all configs are valid",
			server:     New(nil),
			shouldFail: false,
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.shouldFail {
				err = tt.server.Run()
			} else {
				go func() {
					err = tt.server.Run()
				}()
			}
			tt.server.Shutdown()

			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("Run() error = %T, wantErr %T", err, tt.wantErr)
			}
		})
	}
}

func TestServer_Shutdown(t *testing.T) {
	// given
	if err := setupConfig("dev"); err != nil {
		log.Fatalf("Error setting up: %v", err)
	}

	tests := []struct {
		name    string
		server  *Server
		wantErr error
	}{
		{
			name:    "success case: no error returned",
			server:  setupServer(cfg.Env.Server.Host, cfg.Env.Server.Port),
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
