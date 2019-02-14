package service

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/monzo/typhon"
)

// Serve is boilerplate for all main() methods that do nothing but call a service.
func Serve(svc typhon.Service) {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":8000"
	}

	srv, err := typhon.Listen(svc, addr)
	if err != nil {
		panic(err)
	}

	log.Printf("👋  Listening on %v", srv.Listener().Addr())

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	<-done
	log.Printf("☠️  Shutting down")

	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	srv.Stop(c)
}