package server

import (
	"fmt"
	"net/http"

	"github.com/rhynl/http-server/flags"
)

// Start - start the static file server
func Start(args flags.Flags) {
	http.ListenAndServe(fmt.Sprintf("%s:%s", args.Addr, args.Port), server(args))
}

func server(args flags.Flags) http.Handler {
	fmt.Printf("Starting up http-server, serving %s\n", args.Dir)
	fmt.Printf("Available on:\n http://%s:%s\n", args.Addr, args.Port)
	fmt.Println("Hit CTRL-C to stop the server")
	return http.FileServer(http.Dir(args.Dir))
}
