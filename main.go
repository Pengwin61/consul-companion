package main

import (
	"consul-companion/api_consul"
	"consul-companion/internal/cfg"
)

func main() {

	config := cfg.GetConfig()

	// res := consul.GetMembers(config)

	svcList := api_consul.GetNodeServices(config, config.Host)

	for _, r := range svcList.Services {

		api_consul.DeregisterService(config, svcList.Node.Node, r.ID, svcList.Node.Address)
	}

}
