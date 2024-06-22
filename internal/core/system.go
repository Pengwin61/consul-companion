package core

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func gracefulShutdown() {
	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	log.Println("stopping application:", sign)
	os.Exit(0)

}
