package foundation

import (
	"os"
	"sync"

	"github.com/monzo/typhon"
)

type Server struct {
	version string
	name    string
	id      string

	srv *typhon.Server
}

var runOnce sync.Once

// Init sets up and returns a server we can register a service with.
func Init(name string) *Server {
	srv := NewServer(name)
	runOnce.Do(func() {
		//
	})

	return srv
}

func NewServer(name string) *Server {
	srv := &Server{
		version: Version,
		name:    name,
	}

	id, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	srv.id = id

	return srv
}

// ID returns the worker ID of this service, derived from hostname.
func (s *Server) ID() string {
	return s.id
}

// Name returns the name of this service, e.g. service.foo
func (s *Server) Name() string {
	return s.name
}

// Version returns the version of this service, set at compile time from git.
func (s *Server) Version() string {
	return s.version
}

// Run runs a service using httpdown to handle graceful exit.
func (s *Server) Run(svc typhon.Service) {

}
