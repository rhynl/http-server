package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rhynl/http-server/flags"
	"github.com/rhynl/http-server/server"
)

func main() {
	handlerInterrupt()

	args := flags.GetFlags()

	server.Start(args)
}

func handlerInterrupt() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		sig := <-sigs
		if sig == os.Interrupt {
			fmt.Println("http-server stopped.")
			os.Exit(0)
		}
	}()
}
