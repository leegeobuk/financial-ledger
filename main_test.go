package main

import (
	"os"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func Test_getProfile(t *testing.T) {
	tests := []struct {
		name    string
		profile string
		want    string
	}{
		{
			name:    "default case: CONFIG_PROFILE unset",
			profile: "",
			want:    "dev",
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

func Test_initConfig(t *testing.T) {
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
			err := initConfig(tt.profile)
			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
