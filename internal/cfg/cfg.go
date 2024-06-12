package cfg

import (
	"log"
	"os"
)

var Scheme = "http"

type Config struct {
	ConsulAddress string
	ConsulToken   string
}

func GetConfig() Config {
	consulAddress := os.Getenv("CONSUL_HTTP_ADDR")
	if consulAddress == "" {
		log.Fatal("CONSUL_HTTP_ADDR not set")
	}

	consulToken := os.Getenv("CONSUL_HTTP_TOKEN")
	if consulToken == "" {
		log.Fatal("CONSUL_HTTP_TOKEN not set")
	}

	return Config{
		ConsulAddress: consulAddress,
		ConsulToken:   consulToken}

}
