package flags

import (
	"flag"
	"os"
)

// Flags struct
type Flags struct {
	Port string
	Addr string
	Dir  string
}

var argsv []string

const defaultPort = "8080"
const defaultAddr = "0.0.0.0"
const defaultPath = "./"

// GetFlags parse the flags passed to the program
func GetFlags() (args Flags) {
	port := flag.String("p", defaultPort, "Port to use")
	addr := flag.String("a", defaultAddr, "Address to use")
	path := flag.String("path", defaultPath, "directory to be serve")

	a := os.Args[1:]
	if argsv != nil {
		a = argsv
	}

	flag.CommandLine.Parse(a)

	if *path == "" {
		args.Dir = defaultPath
	} else {
		args.Dir = *path
	}

	if *port == "" {
		args.Port = defaultPort
	} else {
		args.Port = *port
	}

	if *addr == "" {
		args.Addr = defaultAddr
	} else {
		args.Addr = *addr
	}

	return
}
