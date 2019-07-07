package backfill

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/gammazero/workerpool"
	"github.com/gosuri/uiprogress"
	"github.com/hokaccha/go-prettyjson"
	"github.com/logrusorgru/aurora"
	"github.com/lolibrary/lolibrary/libraries/filters"
	"github.com/lolibrary/lolibrary/libraries/portforward"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

var (
	errorCount int64 = 0
	errors     []*terrors.Error
	errorsM    sync.Mutex
)

var (
	cmd    io.Closer
	bar    *uiprogress.Bar
	client typhon.Service
	pool   *workerpool.WorkerPool
)

var re = regexp.MustCompile(`^bad_request\.bad_param\.(.*)\.unique`)

func Start(recordType string, count int) {
	fmt.Printf("⚙️  Backfilling %v %s records via the API.\n", aurora.Green(count), recordType)

	concurrency, _ := strconv.ParseInt(os.Getenv("CONCURRENCY"), 10, 64)
	if concurrency <= 0 {
		concurrency = 8
	}

	pool = workerpool.New(int(concurrency))

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
	// wait for everything to complete first.
	pool.StopWait()

	uiprogress.Stop()

	if cmd != nil {
		cmd.Close()
	}

	count := atomic.LoadInt64(&errorCount)

	if count > 0 {
		fmt.Printf("⚠️  Completed with %v errors\n", aurora.Red(count))

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
	pool.Submit(func() {
		request := request

		rsp := request.SendVia(client).Response()

		bar.Incr()

		if rsp.Error != nil {
			if terr, ok := rsp.Error.(*terrors.Error); ok {
				if re.MatchString(terr.Code) {
					// we ignore records that already exist.
					return
				}

				errorsM.Lock()
				errors = append(errors, terr)
				errorsM.Unlock()
			}


			atomic.AddInt64(&errorCount, 1)
		}
	})
}
