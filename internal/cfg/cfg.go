package cfg

import (
	"fmt"
	"log"
	"os"
)

func GetConfig() Config {
	consulAddress := os.Getenv("CONSUL_HTTP_ADDR")
	if consulAddress == "" {
		log.Fatal("CONSUL_HTTP_ADDR not set")
	}

	consulToken := os.Getenv("CONSUL_HTTP_TOKEN")
	if consulToken == "" {
		log.Fatal("CONSUL_HTTP_TOKEN not set")
	}

	consulHTTPScheme := os.Getenv("CONSUL_HTTP_SCHEME")
	if consulHTTPScheme == "" {
		consulHTTPScheme = "http"
		fmt.Println(fmt.Sprintf("Consul Connection URL: %s://%s", consulHTTPScheme, consulAddress))
	} else {
		consulHTTPScheme = "https"
		fmt.Println(fmt.Sprintf("Consul Connection URL: %s://%s", consulHTTPScheme, consulAddress))
	}

	return Config{
		ConsulAddress: consulAddress,
		ConsulToken:   consulToken,
		ConsulScheme:  consulHTTPScheme,
	}
}
