package main

import (
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
			os.Setenv("CONFIG_PROFILE", tt.profile)

			if got := getProfile(); got != tt.want {
				t.Errorf("getProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadConfig(t *testing.T) {
	tests := []struct {
		name    string
		profile string
		wantErr error
	}{
		{
			name:    "fail case: profile=unknown",
			profile: "unknown",
			wantErr: viper.ConfigFileNotFoundError{},
		},
		{
			name:    "success case: profile=local",
			profile: "local",
			wantErr: nil,
		},
		{
			name:    "success case: profile=dev",
			profile: "dev",
			wantErr: nil,
		},
		{
			name:    "success case: profile=stg",
			profile: "stg",
			wantErr: nil,
		},
		{
			name:    "success case: profile=prd",
			profile: "prd",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			err := loadConfig(tt.profile)

			// then
			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("loadConfig() error = %T, wantErr %T", err, tt.wantErr)
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
