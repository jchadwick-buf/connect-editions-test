package main

import (
	"net/http"

	"github.com/jchadwick-buf/connect-editions-test/gen/go/test/testconnect"
	"github.com/jchadwick-buf/connect-editions-test/server"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle(testconnect.NewExampleServiceHandler(server.New()))
	http.ListenAndServe("127.0.0.1:8080", mux)
}
