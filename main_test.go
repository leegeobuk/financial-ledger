package main

import (
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_getProfile(t *testing.T) {
	tests := []struct {
		name    string
		profile string
		want    string
	}{
		{
			name:    "success case: CONFIG_PROFILE unset",
			profile: "",
			want:    "local",
		},
		{
			name:    "success case: CONFIG_PROFILE=local",
			profile: "local",
			want:    "local",
		},
		{
			name:    "success case: CONFIG_PROFILE=dev",
			profile: "dev",
			want:    "dev",
		},
		{
			name:    "success case: CONFIG_PROFILE=stg",
			profile: "stg",
			want:    "stg",
		},
		{
			name:    "success case: CONFIG_PROFILE=prd",
			profile: "prd",
			want:    "prd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("CONFIG_PROFILE", tt.profile)

			if got := getProfile(); got != tt.want {
				t.Errorf("getProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadConfig(t *testing.T) {
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
			err := loadConfig(tt.profile)

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

func Test_setGinMode(t *testing.T) {
	tests := []struct {
		name    string
		profile string
		want    string
	}{
		{
			name:    "success case: profile=unknown",
			profile: "unknown",
			want:    "debug",
		},
		{
			name:    "success case: profile=local",
			profile: "local",
			want:    "debug",
		},
		{
			name:    "success case: profile=dev",
			profile: "dev",
			want:    "test",
		},
		{
			name:    "success case: profile=stg",
			profile: "stg",
			want:    "test",
		},
		{
			name:    "success case: profile=prd",
			profile: "prd",
			want:    "release",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setGinMode(tt.profile)
			if got := gin.Mode(); got != tt.want {
				t.Errorf("setGinMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
