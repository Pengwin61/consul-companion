package main

import (
	"consul-companion/internal/core"
	"flag"
	"strings"
)

func main() {

	// config := cfg.GetConfig()

	pathSearch := flag.String("search", "/opt", "path to search project ")
	flag.Parse()

	core.Path = *pathSearch
	// // res := consul.GetMembers(config)

	// svcList := consul.GetNodeServices(config, config.Host)

	// for _, r := range svcList.Services {

	// 	consul.DeregisterService(config, svcList.Node.Node, r.ID)
	// 	fmt.Println("Deregistered service:", r.ID, "Node:", svcList.Node.Node, "Address:", svcList.Node.Address)
	// }

	prjs, _ := core.GetProjects()
	p := core.GetEnv(prjs)

	// for _, prj := range prjs {
	// 	res := core.GetEnv(prj)
	// 	p = append(p, res)
	// }

	// fmt.Println(p)
	var services []core.ServiceData

	for _, prj := range p {
		for _, env := range prj.Env {
			if strings.Contains(env.Key, "#") {
				continue
			}

			if strings.Contains(env.Key, "EXT") {
				svcName := strings.Replace(env.Key, "EXT_", "", 1)
				svcName = strings.Replace(svcName, "_PORT", "", 1)
				svcName = strings.ToLower(svcName)

				parts := strings.Split(prj.Name, "-")
				index := len(parts) - 1

				partsPrjName := strings.Split(prj.Name, "-")
				projectName := strings.Join(partsPrjName[:len(parts)-1], "-")

				services = append(services, core.ServiceData{
					// Name:     fmt.Sprintf("%s-%s", svcName, prj.Name),
					Name:     svcName,
					Project:  prj.Name,
					Tags:     []string{svcName, parts[index], projectName},
					Port:     env.Value,
					Interval: "5s",
					Timeout:  "5s",
				})
			}
		}
	}

	for _, service := range services {
		core.CreateServiceFile(service)
	}

}
