package main

import (
	"context"
	"log"
	"net/http"
	"reflect"

	"connectrpc.com/connect"
	"github.com/jchadwick-buf/connect-editions-test/gen/go/test"
	"github.com/jchadwick-buf/connect-editions-test/gen/go/test/testconnect"
	"github.com/jchadwick-buf/connect-editions-test/memhttp"
	"github.com/jchadwick-buf/connect-editions-test/server"
	"google.golang.org/protobuf/encoding/prototext"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle(testconnect.NewExampleServiceHandler(server.New()))
	httpServer := memhttp.NewServer(mux)
	httpClient := httpServer.Client()
	client := testconnect.NewExampleServiceClient(httpClient, httpServer.URL())
	response, err := client.ExampleCall(context.TODO(), connect.NewRequest(&test.ExampleMessage{
		NotUtf8: server.ExpectedNotUTF8,
		Flags:   server.ExpectedFlags,
		Child: &test.ExampleMessage{
			NotUtf8: server.ExpectedNotUTF8,
			Enum:    &server.ExpectedEnum,
		},
		Enum: &server.ExpectedEnum,
	}))
	if err != nil {
		log.Fatalf("Server: %v\n", err)
	}
	log.Printf("Response: %s", prototext.MarshalOptions{}.Format(response.Msg))
	if response.Msg.NotUtf8 != server.ReturnedNotUTF8 {
		log.Fatalf("Client: NotUtf8 expected %q, got %q", server.ReturnedNotUTF8, response.Msg.NotUtf8)
	}
	if !reflect.DeepEqual(response.Msg.Flags, server.ReturnedFlags) {
		log.Fatalf("Client: Flags expected %+v, got %+v", server.ReturnedFlags, response.Msg.Flags)
	}
	if response.Msg.Child == nil {
		log.Fatalf("Client: Child expected not nil")
	}
	if response.Msg.Child.NotUtf8 != server.ReturnedNotUTF8 {
		log.Fatalf("Client: Child.NotUtf8 expected %q, got %q", server.ReturnedNotUTF8, response.Msg.Child.NotUtf8)
	}
	if response.Msg.Child.Enum == nil {
		log.Fatalf("Client: Child.Enum expected not nil")
	}
	if *response.Msg.Child.Enum != server.ReturnedEnum {
		log.Fatalf("Client: Child.Enum expected %v, got %v", int32(server.ReturnedEnum), int32(*response.Msg.Child.Enum))
	}
	if response.Msg.Enum == nil {
		log.Fatalf("Client: Enum expected not nil")
	}
	if *response.Msg.Enum != server.ReturnedEnum {
		log.Fatalf("Client: Enum expected %v, got %v", int32(server.ReturnedEnum), int32(*response.Msg.Enum))
	}
	log.Println("Success")
}
