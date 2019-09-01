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

// GetFlags parse the flags passed to the program
func GetFlags() (args Flags) {
	port := flag.String("p", "8080", "Port to use")
	addr := flag.String("a", "0.0.0.0", "Address to use")
	path := flag.String("path", "./", "directory to be serve")

	flag.Parse()

	if *path == "" {
		args.Dir, _ = os.Getwd()
	} else {
		args.Dir = *path
	}

	if *port == "" {
		args.Port = "8080"
	} else {
		args.Port = *port
	}

	if *addr == "" {
		args.Addr = "0.0.0.0"
	} else {
		args.Addr = *addr
	}

	return
}
