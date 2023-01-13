package cfg

import (
	"crypto"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var Env = Config{}

// Config contains all envs loaded from config file.
type Config struct {
	DB     DB     `mapstructure:"db"`
	Token  Token  `mapstructure:"token"`
	Server Server `mapstructure:"server"`
}

// DB contains database-related envs.
type DB struct {
	Type       string `mapstructure:"type"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	Proto      string `mapstructure:"protocol"`
	Addr       string `mapstructure:"address"`
	Schema     string `mapstructure:"schema"`
	Params     string `mapstructure:"params"`
	Migrations string `mapstructure:"migrations"`
}

// DSN returns Data Source Name for db connection.
func (db DB) DSN() string {
	user, password := db.User, db.Password
	proto, addr := db.Proto, db.Addr
	schema, params := db.Schema, db.Params

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", user, password, proto, addr, schema)
	if params != "" {
		dsn = fmt.Sprintf("%s?%s", dsn, params)
	}

	return dsn
}

// Token contains token-related envs.
type Token struct {
	AccessDuration  int    `mapstructure:"accessDuration"`
	RefreshDuration int    `mapstructure:"refreshDuration"`
	PrivatePEM      string `mapstructure:"privatePem"`
	PublicPEM       string `mapstructure:"publicPem"`
	PrivateKey      crypto.PrivateKey
	PublicKey       crypto.PublicKey
	Issuer          string `mapstructure:"issuer"`
}

// Server contains server-related envs.
type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

// Load loads config file located at given path
// according to given profile.
func Load(path, profile string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	if err := viper.Unmarshal(&Env); err != nil {
		return fmt.Errorf("unmarshal envs to config: %w", err)
	}

	return nil
}

// LoadKeys returns private key and public key loaded from pem files.
func LoadKeys(pvtPath, pubPath string) error {
	privatePEM, err := os.ReadFile(filepath.Clean(pvtPath))
	if err != nil {
		return fmt.Errorf("load private.pem: %w", err)
	}

	privateKey, err := jwt.ParseEdPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("parse private.pem: %w", err)
	}

	publicPEM, err := os.ReadFile(filepath.Clean(pubPath))
	if err != nil {
		return fmt.Errorf("load public.pem: %w", err)
	}

	publicKey, err := jwt.ParseEdPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("parse public.pem: %w", err)
	}

	t := &Env.Token
	t.PrivateKey = privateKey
	t.PublicKey = publicKey

	return nil
}
