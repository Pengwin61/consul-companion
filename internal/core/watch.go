package core

import (
	"log"
	"time"
)

func RunWatch() {
	watch()
}

func watch() {
	isSrarted := true
	errCh := make(chan error)

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

		RunCreatesServices(errCh)
		time.Sleep(60 * time.Second)

	}
}
