package main

import (
	"log"
	"net/http"

	. "github.com/usvc/go-server"
)

func main() {
	// initialisation

	config := NewConfigFromEnvironment()
	server := NewServer(config)
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	server.SetHandler(handler)

	// middleware hook 1: via callback

	config.OnError = func(err error) {
		log.Printf("[from callback] server error: %s", err)
	}
	config.OnListening = func() {
		log.Printf("[from callback] listening at %s", config.GetURL())
	}
	config.OnListeningTimeout = func() {
		log.Println("[from callback] failed to get a response from server")
	}
	config.OnShutdown = func() {
		log.Println("[from callback] server is shutting down")
	}
	config.OnStarting = func() {
		log.Println("[from callback] server is starting up")
	}

	// middleware hook 2: via channel

	go func(serverSignal <-chan Signal) {
		for {
			select {
			case signal := <-serverSignal:
				switch signal.GetCode() {
				case SignalError:
					log.Printf("[from channel] server error: %s", signal.GetError())
				case SignalListening:
					log.Printf("[from channel] listening at %s", config.GetURL())
				case SignalListeningTimeout:
					log.Println("[from channel] failed to get a response from server")
				case SignalShutdown:
					log.Println("[from channel] server is shutting down")
				case SignalStarting:
					log.Println("[from channel] server is starting up")
				}
			}
		}
	}(server.GetSignals())

	// start the server

	server.Start()
}
