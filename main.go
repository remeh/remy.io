// Quick'n'basic web front stuff
// for remy.io
package main

import (
	"net/http"
)

type CvHandler struct {
}

func (c *CvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(302)
	w.Header().Set("Location", "cv.html")
}

func main() {
	// CV redirects of old uris.
	cv := CvHandler{}
	http.Handle("/cv", &cv)
	http.Handle("/resume", &cv)

	// And static serving.
	http.Handle("/", http.FileServer(http.Dir("web")))

	http.ListenAndServe(":8080", nil)
}
