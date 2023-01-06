package cfg

var Env = Config{}

// Config contains all envs loaded from config file.
type Config struct {
	DB     DB     `yaml:"db"`
	Server Server `yaml:"server"`
}

// DB contains database-related envs.
type DB struct {
	DBType   string `yaml:"type"`
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Server contains server-related envs.
type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
