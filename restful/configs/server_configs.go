package configs

type ServerConfiguration struct {
	APIPort    int
	MetricPort int
}

func NewServerConfiguration() *ServerConfiguration {
	return &ServerConfiguration{
		APIPort:    8180,
		MetricPort: 9090,
	}
}
