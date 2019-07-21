package foundation

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/lolibrary/lolibrary/libraries/filters"
	"github.com/lolibrary/lolibrary/libraries/logging"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

type Server struct {
	version string
	name    string
	id      string

	srv *typhon.Server
}

func init() {
	typhon.Client = typhon.Client.Filter(filters.ClientFilters)
}

var runOnce sync.Once

// Init sets up and returns a server we can register a service with.
func Init(name string) *Server {
	srv := NewServer(name)
	runOnce.Do(func() {
		// set up logging
		// logging.Init(name)
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

// Run runs a service with typhon, handling shutdown gracefully.
func (s *Server) Run(svc typhon.Service) {
	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	started := make(chan error)
	go func() {
		started <- s.Start(svc)
	}()

	select {
	case err := <-started:
		if err != nil {
			panic(err)
		}
	case s := <-shutdown:
		slog.Info(nil, "⚠️  Got a shutdown signal (%v) before service started, exiting", s)
		os.Exit(0)
	case <-time.After(time.Minute * 5):
		slog.Critical(nil, "😨 didn't start %s (worker: %s) in time (5m)", s.Name(), s.ID())
		os.Exit(1)
	}

	select {
	case s := <-shutdown:
		slog.Info(nil, "✋  Got a %v, shutting down", s)
	case err := <-s.srv.Done():
		slog.Error(nil, "🛑  Typhon died on us, shutting down: %v", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	s.Stop(ctx)
	slog.Info(nil, "👋  Server stopped")
}

func (s *Server) Stop(ctx context.Context) {
	done := make(chan bool)
	go func() {
		s.srv.Stop(ctx)

		done <- true
	}()

	select {
	case <-done:
	case <-ctx.Done():
	}
}

func (s *Server) Start(svc typhon.Service) error {
	fmt.Fprintf(os.Stderr, "🦔 💨  %s; version %s | %s, %s/%s\n",
		s.Name(), Version, runtime.Version(), runtime.GOOS, runtime.GOARCH,
	)

	srv, err := typhon.Listen(svc, "")
	if err != nil {
		return err
	}

	slog.Info(nil, "👂 %s listening on %v", s.Name(), srv.Listener().Addr())
	s.srv = srv

	return nil
}
