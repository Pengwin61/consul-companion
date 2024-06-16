package main

import (
	"consul-companion/internal/core"
	"flag"
)

func init() {
	pathSearch := flag.String("search", "/opt", "path to search project ")
	flag.Parse()

	core.Path = *pathSearch
}

func main() {

	// config := cfg.GetConfig()

	// // res := consul.GetMembers(config)

	// svcList := consul.GetNodeServices(config, config.Host)

	// for _, r := range svcList.Services {

	// 	consul.DeregisterService(config, svcList.Node.Node, r.ID)
	// 	fmt.Println("Deregistered service:", r.ID, "Node:", svcList.Node.Node, "Address:", svcList.Node.Address)
	// }

	core.RunCreatesServices()

}
