// basic webserver with http2 cleartext support
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	addr := flag.String("addr", ":8080", "tcp address to listen on for incoming requests")
	flag.Parse()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprint(w, string(requestDump))
	})

	h2s := &http2.Server{
		// ...
	}
	h1s := &http.Server{
		Addr:    *addr,
		Handler: h2c.NewHandler(handler, h2s),
	}
	log.Println("Listening on", *addr)
	log.Fatal(h1s.ListenAndServe())
}
