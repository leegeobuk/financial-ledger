package cfg

import "fmt"

var Env = Config{}

// Config contains all envs loaded from config file.
type Config struct {
	DB     DB     `yaml:"db"`
	Server Server `yaml:"server"`
}

// DB contains database-related envs.
type DB struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Addr     string `yaml:"address"`
	Schema   string `yaml:"schema"`
}

// DSN returns Data Source Name for db connection.
func (db DB) DSN() string {
	return fmt.Sprintf("%s:%s@%s/%s", db.User, db.Password, db.Addr, db.Schema)
}

// Server contains server-related envs.
type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
