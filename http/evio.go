package main

import (
	"github.com/tidwall/evio"
)

func main() {
	responseString := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 12\r\n\r\nHello World\r\n"
	responseBuffer := []byte(responseString)

	var events evio.Events
	events.Data = func(id int, in []byte) (out []byte, action evio.Action) {
		// fmt.Println(len(in))
		numberOfRequests := len(in) / 40

		for i := 0; i < numberOfRequests; i++ {
			out = append(out, responseBuffer...)
		}
		return
	}
	if err := evio.Serve(events, "tcp://0.0.0.0:8000"); err != nil {
		panic(err.Error())
	}
}
