package main

import (
	"consul-companion/internal/cfg"
	"consul-companion/internal/consul"
)

func main() {

	config := cfg.GetConfig()

	// res := consul.GetMembers(config)

	response := consul.GetNodeServices(config, "gitlab-runner")

	for _, r := range response.Services {

		consul.DeregisterService(config, response.Node.Node, r.ID)
	}

}
