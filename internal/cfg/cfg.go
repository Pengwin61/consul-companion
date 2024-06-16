package cfg

import (
	"consul-companion/internal/core"
	"flag"
	"fmt"
	"log"
	"os"
)

var ConsulHTTPScheme string

type Config struct {
	ConsulAddress string
	ConsulToken   string
	Host          string
	Path          string
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

	ConsulHTTPScheme = os.Getenv("CONSUL_HTTP_SCHEME")
	if ConsulHTTPScheme == "" {
		ConsulHTTPScheme = "http"
		fmt.Println(fmt.Sprintf("Consul Connection URL: %s://%s", ConsulHTTPScheme, consulAddress))

	}

	// Получаем хост из флага
	host := flag.String("host", "127.0.0.1", "hosts ")
	// if *host == "" {
	// 	log.Fatal("host is empty, i can`t get host")
	// }

	path := flag.String("path", "/etc/consul/consul.d", "path to config ")
	if *path == "" {
		log.Fatal("path is empty, i can`t get path")
	}

	pathSearch := flag.String("search", "/opt", "path to search project ")
	core.Path = *pathSearch

	flag.Parse()

	return Config{
		ConsulAddress: consulAddress,
		ConsulToken:   consulToken,
		Host:          *host,
		Path:          *path,
	}

}
