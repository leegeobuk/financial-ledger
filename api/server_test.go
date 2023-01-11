package api

import (
	"reflect"
	"strings"
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
		shouldFail bool
		server     *Server
		wantErrStr string
	}{
		{
			name:       "fail case: port=xxx",
			shouldFail: true,
			server:     setupServer(host, "xxx"),
			wantErrStr: "run api server",
		},
		{
			name:       "success case: all configs are valid",
			shouldFail: false,
			server:     New(nil),
			wantErrStr: "",
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
			_ = tt.server.Shutdown()

			// then
			if tt.shouldFail {
				if got := err.Error(); !strings.HasPrefix(got, tt.wantErrStr) {
					t.Errorf("Run() error = %v, wantErrStr %s", err, tt.wantErrStr)
				}

				return
			}

			if err != nil {
				t.Errorf("Run() error = %v, wantErrStr %s", err, tt.wantErrStr)
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
				_ = tt.server.Run()
			}()

			err := tt.server.Shutdown()

			// then
			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("Shutdown() error = %v, wantErrStr %v", err, tt.wantErr)
			}
		})
	}
}
