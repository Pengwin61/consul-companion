package main

import (
	"consul-companion/internal/cfg"
	"consul-companion/internal/consul"
)

func main() {

	config := cfg.GetConfig()

	// res := consul.GetMembers(config)

	svcList := consul.GetNodeServices(config, config.Host)

	for _, r := range svcList.Services {

		consul.DeregisterService(config, svcList.Node.Node, r.ID, svcList.Node.Address)
	}

}
