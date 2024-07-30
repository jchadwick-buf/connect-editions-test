package server

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"connectrpc.com/connect"
	"github.com/jchadwick-buf/connect-editions-test/gen/go/test"
	"github.com/jchadwick-buf/connect-editions-test/gen/go/test/testconnect"
	"google.golang.org/protobuf/encoding/prototext"
)

const (
	ExpectedNotUTF8 = "te\x00st1"
	ReturnedNotUTF8 = "te\x00st2"
)

var (
	ExpectedEnum  = test.ExampleEnum_VALUE_IN
	ReturnedEnum  = test.ExampleEnum_VALUE_OUT
	ExpectedFlags = []bool{false, true}
	ReturnedFlags = []bool{true, false}
)

func New() testconnect.ExampleServiceHandler {
	return &svc{}
}

type svc struct{}

func (svc) ExampleCall(ctx context.Context, request *connect.Request[test.ExampleMessage]) (*connect.Response[test.ExampleMessage], error) {
	log.Printf("Request: %s", prototext.MarshalOptions{}.Format(request.Msg))
	if request.Msg.NotUtf8 != ExpectedNotUTF8 {
		return nil, fmt.Errorf("NotUtf8 expected %q, got %q", ExpectedNotUTF8, request.Msg.NotUtf8)
	}
	if !reflect.DeepEqual(request.Msg.Flags, ExpectedFlags) {
		return nil, fmt.Errorf("Flags expected %+v, got %+v", ExpectedFlags, request.Msg.Flags)
	}
	if request.Msg.Child == nil {
		return nil, fmt.Errorf("Child expected not nil")
	}
	if request.Msg.Child.NotUtf8 != ExpectedNotUTF8 {
		return nil, fmt.Errorf("Child.NotUtf8 expected %q, got %q", ExpectedNotUTF8, request.Msg.Child.NotUtf8)
	}
	if request.Msg.Child.Enum == nil {
		return nil, fmt.Errorf("Child.Enum expected not nil")
	}
	if *request.Msg.Child.Enum != ExpectedEnum {
		return nil, fmt.Errorf("Child.Enum expected %v, got %v", int32(ExpectedEnum), int32(*request.Msg.Child.Enum))
	}
	if request.Msg.Enum == nil {
		return nil, fmt.Errorf("Enum expected not nil")
	}
	if *request.Msg.Enum != ExpectedEnum {
		return nil, fmt.Errorf("Enum expected %v, got %v", int32(ExpectedEnum), int32(*request.Msg.Enum))
	}
	return connect.NewResponse(&test.ExampleMessage{
		NotUtf8: ReturnedNotUTF8,
		Flags:   ReturnedFlags,
		Child: &test.ExampleMessage{
			NotUtf8: ReturnedNotUTF8,
			Enum:    &ReturnedEnum,
		},
		Enum: &ReturnedEnum,
	}), nil
}
