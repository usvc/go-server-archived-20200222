package main

import (
	"crypto/tls"
	"net/http"
	"sync"
	"time"
)

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
		instance: &http.Server{
			Addr: config.GetAddr(),
		},
		signals: make(chan Signal),
	}
}

type Server struct {
	config       *Config
	err          error
	instance     *http.Server
	isResponding bool
	signals      chan Signal
	waiter       sync.WaitGroup
}

func (server *Server) GetSignals() <-chan Signal {
	return server.signals
}

func (server *Server) SetHandler(handler http.Handler) {
	server.instance.Handler = handler
}

func (server *Server) StartWithHandler(handler http.Handler) {
	server.SetHandler(handler)
	server.Start()
}

func (server *Server) Start() {
	if server.config.OnShutdown != nil {
		server.instance.RegisterOnShutdown(server.config.OnShutdown)
	}
	server.instance.RegisterOnShutdown(server.handleShutdown)
	server.waiter.Add(1)
	server.handleStarting()
	go server.listenAndServe()
	every := time.Tick(server.config.LivenessCheckInterval)
	until := time.After(server.config.LivenessCheckTimeout)
	go server.monitorServer(every, until)
	server.waiter.Wait()
}

// isResponsive checks if the server instance is responsive to requests
// and returns a true on status code match with :expectedStatusCode
func (server *Server) checkIfResponsive() bool {
	request, err := http.NewRequest(server.config.LivenessCheckMethod, server.config.GetURL()+server.config.LivenessCheckPath, nil)
	if err != nil {
		return false
	}
	client := &http.Client{}
	if server.config.HasValidTLS() {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
	response, err := client.Do(request)
	if err != nil {
		return false
	}
	return response.StatusCode == server.config.LivenessCheckStatusCode
}

func (server *Server) handleError(err error) {
	server.err = err
	if server.config.OnError != nil {
		server.config.OnError(err)
	}
	server.signals <- ServerSignal{SignalError, err}
}

func (server *Server) handleListening() {
	if server.config.OnListening != nil {
		server.config.OnListening()
	}
	server.signals <- ServerSignal{SignalListening, nil}
}

func (server *Server) handleListeningTimeout() {
	if server.config.OnListeningTimeout != nil {
		server.config.OnListeningTimeout()
	}
	server.signals <- ServerSignal{SignalListeningTimeout, nil}
}

func (server *Server) handleShutdown() {
	if server.config.OnShutdown != nil {
		server.config.OnShutdown()
	}
	server.signals <- ServerSignal{SignalShutdown, nil}
}

func (server *Server) handleStarting() {
	if server.config.OnStarting != nil {
		server.config.OnStarting()
	}
	server.signals <- ServerSignal{SignalStarting, nil}
}

func (server *Server) listenAndServe() {
	var httpListenAndServeError error
	if server.config.HasValidTLS() {
		httpListenAndServeError = server.instance.ListenAndServeTLS(
			server.config.TLSCertificatePath,
			server.config.TLSKeyPath,
		)
	} else {
		httpListenAndServeError = server.instance.ListenAndServe()
	}
	if httpListenAndServeError != nil {
		server.handleError(httpListenAndServeError)
	}
	server.waiter.Done()
}

func (server *Server) monitorServer(checkEvery <-chan time.Time, checkUntil <-chan time.Time) {
	hasTimedOut := false
	for {
		if hasTimedOut {
			server.handleListeningTimeout()
			break
		} else if server.isResponding {
			server.handleListening()
			break
		}
		select {
		case <-checkEvery:
			server.isResponding = server.checkIfResponsive()
			if !server.isResponding {
				server.signals <- ServerSignal{SignalNotLiveYet, nil}
			}
		case <-checkUntil:
			hasTimedOut = true
		}
	}
}
