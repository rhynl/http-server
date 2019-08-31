package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type flags struct {
	port string
	addr string
	dir  string
}

func main() {

	args := getFlags()

	http.ListenAndServe(fmt.Sprintf("%s:%s", args.addr, args.port), http.FileServer(http.Dir(args.dir)))
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
