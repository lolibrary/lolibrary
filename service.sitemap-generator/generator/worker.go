package generator

import (
	"context"
	"time"

	"github.com/monzo/slog"
)

var (
	stop     = make(chan bool)
	finished = make(chan bool)
)

func Start() {
	go func() {
		ctx := context.Background()
		for {
			select {
			case <-stop:
				finished <- true
				return
			default:
			}

			if err := Generate(ctx); err != nil {
				slog.Error(ctx, "Error generating sitemap: %v", err)
				// todo: track error
			}

			<-time.After(time.Minute * 30)
		}
	}()
}

func Stop() {
	stop <- true
	select {
	case <-finished:
		close(finished)
	case <-time.After(time.Second * 5):
	}

	close(stop)
}
