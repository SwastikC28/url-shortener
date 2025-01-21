package httpserver

import "github.com/ilyakaznacheev/cleanenv"

type httpserverconfig struct {
	Port int `env:"API-PORT" env-default:"8080"`
}

func newhttpserverconfig() *httpserverconfig {
	config := httpserverconfig{}
	cleanenv.ReadEnv(&config)

	return &config
}
