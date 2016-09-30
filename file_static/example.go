package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var port = flag.String("p", "8888", "static file server port")

func main() {

	flag.Parse()

	server_port := *port

	println("start static file server :" + server_port)

	mux := http.NewServeMux()
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))
	log.Fatal(http.ListenAndServe(":"+server_port, mux))
}
