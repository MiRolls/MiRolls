package config

type Config struct {
	Server   *Server   `yaml:"server"`
	Database *Database `yaml:"database"`
	Site     string
}

type Server struct {
	Port   int    `yaml:"port"`
	Static string `yaml:"static"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}
