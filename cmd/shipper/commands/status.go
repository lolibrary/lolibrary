package commands

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/gammazero/workerpool"
	"github.com/logrusorgru/aurora"
	"github.com/lolibrary/lolibrary/libraries/filters"
	"github.com/lolibrary/lolibrary/libraries/portforward"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
	"gopkg.in/urfave/cli.v1"
)

var services = []string{
	"service.brand",
	"service.category",
	"service.feature",
	"service.color",
	"service.tag",
	"service.attribute",
	"service.item",
	"service.image",
	"service.sitemap-generator",
}

type versionResponse struct {
	Version string `json:"version"`
}

var (
	mutex    sync.Mutex
	versions = map[string]string{
		"service.brand":             "unknown",
		"service.category":          "unknown",
		"service.feature":           "unknown",
		"service.color":             "unknown",
		"service.tag":               "unknown",
		"service.attribute":         "unknown",
		"service.item":              "unknown",
		"service.image":             "unknown",
		"service.sitemap-generator": "unknown",
	}
)

func Status(_ *cli.Context) {
	port, forward := portforward.Enable()
	defer forward.Close()

	client := typhon.Client.
		Filter(filters.EdgeProxyFilter(fmt.Sprintf("127.0.0.1:%d", port))).
		Filter(typhon.ErrorFilter)

	pool := workerpool.New(8)
	ctx := context.Background()

	fmt.Println("⏳ Waiting for version responses...")

	for _, s := range services {
		s := s
		pool.Submit(func() {
			var version versionResponse

			req := typhon.NewRequest(ctx, "GET", fmt.Sprintf("http://%s/foundation/version", s), nil)
			rsp := req.SendVia(client).Response()

			mutex.Lock()
			defer mutex.Unlock()

			if rsp.Error != nil {
				versions[s] = rsp.Error.(*terrors.Error).Code

				return
			}

			if err := rsp.Decode(&version); err != nil {
				versions[s] = err.Error()

				return
			}

			if version.Version == "" {
				fmt.Println("version was empty")
				os.Exit(1)
			}

			versions[s] = version.Version
		})
	}

	pool.StopWait()

	for _, service := range services {
		v := versions[service]
		svc := service + strings.Repeat(" ", 30)

		var val aurora.Value
		switch {
		case v == "unknown":
			val = aurora.Blue(v)
		case len(v) == 40:
			val = aurora.Green(v)
		default:
			val = aurora.Red(v)
		}

		fmt.Printf("  %v :  %v\n", aurora.Blue(svc[:30]), val)
	}
}
