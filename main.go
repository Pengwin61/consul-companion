package main

import (
	"consul-companion/internal/cfg"
	"consul-companion/internal/consul"
	"fmt"
)

func main() {

	config := cfg.GetConfig()

	// res := consul.GetMembers(config)

	svcList := consul.GetNodeServices(config, config.Host)

	for _, r := range svcList.Services {

		// consul.DeregisterService(config, svcList.Node.Node, r.ID)
		fmt.Println("Deregistered service:", r.ID, "Node:", svcList.Node.Node, "Address:", svcList.Node.Address)
	}

}
