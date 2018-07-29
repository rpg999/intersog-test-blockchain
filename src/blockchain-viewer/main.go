package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	// Listening for tcp connection on on port 6868
	address := "127.0.0.1:6868"
	ln, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatalf("unable to listen to the address %a: %s", address, err)
	}

	// Create server with some config
	var mux = newServerMux()
	server := &http.Server{
		Handler:      mux,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	// Start the server
	log.Printf("Listen for requests at %q, press Ctrl+C to stop webapp server", address)
	if err := server.Serve(ln); err != nil {
		log.Fatalf("HTTP fail serve on %q: %s", address, err)
	}
}
