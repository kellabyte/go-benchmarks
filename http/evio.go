package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/tidwall/evio"
)

func main() {
	logging := false
	if len(os.Args) > 1 {
		if os.Args[1] == "--debug" {
			logging = true
		}
	}

	const responseString = "HTTP/1.1 204\r\n\r\n"
	responseBuffer := []byte(responseString)

	buffer := make([]byte, 0, len(responseBuffer)*10000)
	for i := 0; i < 10000; i++ {
		buffer = append(buffer, responseBuffer...)
	}

	var events evio.Events
	events.Serving = func(s evio.Server) (action evio.Action) {
		log.Printf("serving %s\n", s.Addrs[0])
		return
	}
	events.Opened = func(id int, info evio.Info) (out []byte, opts evio.Options, ctx interface{}, action evio.Action) {
		// Reuse the input buffer on Data callbacks to avoid unneeded alloc.
		opts.ReuseInputBuffer = true
		return
	}
	events.Data = func(id int, ctx interface{}, in []byte) (out []byte, action evio.Action) {
		if logging {
			fmt.Printf("%q\n", in)
		}
		numberOfRequests := len(in) / 40
		out = buffer[:numberOfRequests*len(responseBuffer)]
		return
	}
	events.NumLoops = runtime.NumCPU() / 2
	evio.Serve(events, "tcp://0.0.0.0:8000?reuseport=true")
}
