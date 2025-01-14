package server

type ServerList struct {
	App            App
	Grpc           Server
	UserManagement Server
	EMoney         Server
	TopUP          Server
}

type Server struct {
	TLS     bool   `yaml:"tls"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Timeout int    `yaml:"timeout"`
}

type App struct {
	Name string `yaml:"Name"`
}
