package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var dir string

	port := flag.String("port", "3001", "port in which the server will be listen on.")
	path := flag.String("path", "", "directory to be serve")

	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}

	http.ListenAndServe(fmt.Sprintf(":%s", *port), http.FileServer(http.Dir(dir)))
}
