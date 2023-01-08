package cfg

import "fmt"

var Env = Config{}

// Config contains all envs loaded from config file.
type Config struct {
	DB     DB     `mapstructure:"db"`
	Server Server `mapstructure:"server"`
}

// DB contains database-related envs.
type DB struct {
	Type     string `mapstructure:"type"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Proto    string `mapstructure:"protocol"`
	Addr     string `mapstructure:"address"`
	Schema   string `mapstructure:"schema"`
}

// DSN returns Data Source Name for db connection.
func (db DB) DSN() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", db.User, db.Password, db.Proto, db.Addr, db.Schema)
}

// Server contains server-related envs.
type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
