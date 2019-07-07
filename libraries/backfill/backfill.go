package backfill

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gosuri/uiprogress"
	"github.com/hokaccha/go-prettyjson"
	"github.com/logrusorgru/aurora"
	"github.com/lolibrary/lolibrary/libraries/filters"
	"github.com/lolibrary/lolibrary/libraries/portforward"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

var (
	errorCount = 0
	errors     []*terrors.Error
)

var (
	cmd    io.Closer
	bar    *uiprogress.Bar
	client typhon.Service
)

func Start(recordType string, count int) {
	fmt.Printf("⚙️  Backfilling %v %s records via the API.\n", aurora.Green(count), recordType)

	port, closer := portforward.Enable()
	cmd = closer

	client = typhon.Client.
		Filter(filters.EdgeProxyFilter(fmt.Sprintf("127.0.0.1:%d", port))).
		Filter(typhon.ErrorFilter)

	uiprogress.Start()
	bar = uiprogress.AddBar(count)
	bar.AppendCompleted()
}

func Stop() {
	uiprogress.Stop()

	if cmd != nil {
		cmd.Close()
	}

	if errorCount > 0 {
		fmt.Printf("⚠️  Completed with %v errors\n", aurora.Red(errorCount))

		if debug, _ := strconv.ParseBool(os.Getenv("DEBUG")); debug {
			for _, terr := range errors {
				json := map[string]interface{}{
					"code":    terr.Code,
					"message": terr.Message,
				}
				if len(terr.Params) > 0 {
					json["params"] = terr.Params
				}

				b, err := prettyjson.Marshal(json)
				if err != nil {
					fmt.Printf("%v: failed to marshal terror to json: %v\n", aurora.Red("error").Bold(), err)
					continue
				}

				fmt.Println(string(b))
				fmt.Printf("\n")
			}
		}
	} else {
		fmt.Println("✅ All done!")
	}
}

func Request(request typhon.Request) {
	rsp := request.SendVia(client).Response()
	if rsp.Error != nil {
		errorCount++

		if terr, ok := rsp.Error.(*terrors.Error); ok {
			errors = append(errors, terr)
		}
	}

	bar.Incr()
}
