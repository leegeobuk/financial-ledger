package api

import (
	"net"
	"reflect"
	"testing"

	"github.com/leegeobuk/financial-ledger/cfg"
	"github.com/leegeobuk/financial-ledger/testutil"
)

func TestServer_Run(t *testing.T) {
	// given
	if err := testutil.SetupConfig("local"); err != nil {
		t.Fatalf("Error setting up config: %v", err)
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
			// when
			var err error
			if tt.shouldFail {
				err = tt.server.Run()
			} else {
				go func() {
					err = tt.server.Run()
				}()
			}
			tt.server.Shutdown()

			// then
			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("Run() error = %T, wantErr %T", err, tt.wantErr)
			}
		})
	}
}

func TestServer_Shutdown(t *testing.T) {
	// given
	if err := testutil.SetupConfig("local"); err != nil {
		t.Fatalf("Error setting up config: %v", err)
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
			// when
			go func() {
				tt.server.Run()
			}()

			err := tt.server.Shutdown()

			// then
			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
