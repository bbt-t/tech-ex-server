package apiserver

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ExitHandler() {
	signalExit := make(chan os.Signal, 1)
	signal.Notify(signalExit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			s := <-signalExit
			switch s {
			case os.Interrupt:
				fallthrough
			case syscall.SIGINT:
				fallthrough
			case syscall.SIGTERM:
				msg := "---> SERVER STOPPED! <---"
				log.Fatal(msg)
			}
		}
	}()
}
