package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	const port = "8080"

	var rootDirPtr = flag.String("rootDir", "/root/gowebserver", "Use to change working directory")
	flag.Parse()

	applyHeaders := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
			w.Header().Set("Pragma", "no-cache")

			next.ServeHTTP(w, r)
		})
	}

	// handlers
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(*rootDirPtr+"/public")))
	mux.HandleFunc("/test", test)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/testt", http.NotFoundHandler())

	fmt.Println("Listening on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, applyHeaders(mux)))
}

func test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Test</h1>")
}
