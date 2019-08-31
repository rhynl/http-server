package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type flags struct {
	port string
	addr string
	dir  string
}

func main() {
	handlerInterrupt()

	args := getFlags()

	http.ListenAndServe(fmt.Sprintf("%s:%s", args.addr, args.port), server(args))
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

func getFlags() (args flags) {
	port := flag.String("p", "8080", "Port to use")
	addr := flag.String("a", "0.0.0.0", "Address to use")
	path := flag.String("path", "./", "directory to be serve")

	flag.Parse()

	if *path == "" {
		args.dir, _ = os.Getwd()
	} else {
		args.dir = *path
	}

	if *port == "" {
		args.port = "8080"
	} else {
		args.port = *port
	}

	if *addr == "" {
		args.addr = "0.0.0.0"
	} else {
		args.addr = *addr
	}

	return
}

func server(args flags) http.Handler {
	fmt.Printf("Starting up http-server, serving %s\n", args.dir)
	fmt.Printf("Available on:\n http://%s:%s\n", args.addr, args.port)
	fmt.Println("Hit CTRL-C to stop the server")
	return http.FileServer(http.Dir(args.dir))
}
