// Quick'n'basic web front stuff
// for remy.io
package main

import (
	"net/http"
	"os"
)

func main() {
	if len(os.Getenv("DOCKER")) > 0 {
		http.Handle("/", http.FileServer(http.Dir("/remy.io/web/")))
	} else {
		http.Handle("/", http.FileServer(http.Dir("web/")))
	}

	http.ListenAndServe(":8080", nil)
}
