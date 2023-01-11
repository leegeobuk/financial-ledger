package cfg

import (
	"strings"
	"testing"
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
			want:    "user:1111@tcp(127.0.0.1:3306)/household_ledger?multiStatements=true",
		},
		{
			name:    "success case: profile=dev",
			profile: "dev",
			want:    "user:1111@tcp(ledger-db:3306)/household_ledger?multiStatements=true",
		},
		{
			name:    "success case: profile=stg",
			profile: "stg",
			want:    "user:1111@tcp(ledger-db:3306)/household_ledger?multiStatements=true",
		},
		{
			name:    "success case: profile=prd",
			profile: "prd",
			want:    "user:1111@tcp(ledger-db:3306)/household_ledger?multiStatements=true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			if err := Load(".", tt.profile); err != nil {
				t.Fatalf("Error setting up config: %v", err)
			}

			if got := Env.DB.DSN(); got != tt.want {
				t.Errorf("DSN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name       string
		shouldFail bool
		profile    string
		wantErrStr string
	}{
		{
			name:       "fail case: profile=unknown",
			shouldFail: true,
			profile:    "unknown",
			wantErrStr: "load config",
		},
		{
			name:       "success case: profile=local",
			shouldFail: false,
			profile:    "local",
			wantErrStr: "",
		},
		{
			name:       "success case: profile=dev",
			shouldFail: false,
			profile:    "dev",
			wantErrStr: "",
		},
		{
			name:       "success case: profile=stg",
			shouldFail: false,
			profile:    "stg",
			wantErrStr: "",
		},
		{
			name:       "success case: profile=prd",
			shouldFail: false,
			profile:    "prd",
			wantErrStr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			err := Load(".", tt.profile)

			// then
			if tt.shouldFail {
				if got := err.Error(); !strings.HasPrefix(got, tt.wantErrStr) {
					t.Errorf("loadConfig() error = %v, wantErrStr %s", err, tt.wantErrStr)
				}

				return
			}

			if err != nil {
				t.Errorf("loadConfig() error = %v, wantErrStr %s", err, tt.wantErrStr)
			}
		})
	}
}
