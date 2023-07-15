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

	mux := http.NewServeMux()

	// handlers
	mux.Handle("/", http.FileServer(http.Dir(*rootDirPtr+"/public")))
	mux.HandleFunc("/test", test)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/testt", http.NotFoundHandler())

	fmt.Println("Listening on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Test</h1>")
}
