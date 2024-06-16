package main

import (
	"consul-companion/internal/core"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	pathSearch := flag.String("search", "/opt", "path to search project ")
	pathConfig := flag.String("path", "/etc/consul/consul.d", "path to config ")
	flag.Parse()

	core.Path = *pathSearch
	core.CFG_PATH = *pathConfig

}

func main() {
	isSrarted := true
	errCh := make(chan error)

	// cfg := cfg.GetConfig()

	// // res := consul.GetMembers(config)

	// svcList := consul.GetNodeServices(config, config.Host)

	// for _, r := range svcList.Services {

	// 	consul.DeregisterService(config, svcList.Node.Node, r.ID)
	// 	fmt.Println("Deregistered service:", r.ID, "Node:", svcList.Node.Node, "Address:", svcList.Node.Address)
	// }

	go func() {
		for err := range errCh {
			log.Println(err)
		}
	}()

	go gracefulShutdown()

	for {
		if isSrarted {
			log.Println("Starting...")
			isSrarted = false
		}

		core.RunCreatesServices(errCh)
		time.Sleep(10 * time.Second)

	}

}

func gracefulShutdown() {
	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	log.Println("stopping application:", sign)
	os.Exit(0)

}
