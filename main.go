// Quick'n'basic web front stuff
// for remy.io
package main

import (
	"net/http"
	"os"
)

func main() {
	if len(os.Getenv("LOCAL")) > 0 {
		http.Handle("/", http.FileServer(http.Dir("web/")))
	} else {
		http.Handle("/", http.FileServer(http.Dir("/home/webapps/live/remy.io/web/")))
	}

	http.ListenAndServe(":8080", nil)
}
