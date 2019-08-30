package main

import (
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()

	http.ListenAndServe(":8080", http.FileServer(http.Dir(dir)))
}
