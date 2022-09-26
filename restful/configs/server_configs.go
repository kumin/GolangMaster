package configs

type ServerConfiguration struct {
	Port int
}

func NewServerConfiguration() *ServerConfiguration {
	return &ServerConfiguration{
		Port: 8080,
	}
}
