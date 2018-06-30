package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var contents []byte
var filename string

func main() {
	flag.StringVar(&filename, "filename", "file", "The filename to read and serve via HTTP.")
	flag.Parse()

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to read given file.", err)
	}

	log.Print("Starting fileserver on :3000")

	http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(contents)
	}))
}
