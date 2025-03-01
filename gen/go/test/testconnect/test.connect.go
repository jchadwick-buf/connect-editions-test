// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: test/test.proto

package testconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	test "github.com/jchadwick-buf/connect-editions-test/gen/go/test"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ExampleServiceName is the fully-qualified name of the ExampleService service.
	ExampleServiceName = "test.ExampleService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ExampleServiceExampleCallProcedure is the fully-qualified name of the ExampleService's
	// ExampleCall RPC.
	ExampleServiceExampleCallProcedure = "/test.ExampleService/ExampleCall"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	exampleServiceServiceDescriptor           = test.File_test_test_proto.Services().ByName("ExampleService")
	exampleServiceExampleCallMethodDescriptor = exampleServiceServiceDescriptor.Methods().ByName("ExampleCall")
)

// ExampleServiceClient is a client for the test.ExampleService service.
type ExampleServiceClient interface {
	ExampleCall(context.Context, *connect.Request[test.ExampleMessage]) (*connect.Response[test.ExampleMessage], error)
}

// NewExampleServiceClient constructs a client for the test.ExampleService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewExampleServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ExampleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &exampleServiceClient{
		exampleCall: connect.NewClient[test.ExampleMessage, test.ExampleMessage](
			httpClient,
			baseURL+ExampleServiceExampleCallProcedure,
			connect.WithSchema(exampleServiceExampleCallMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// exampleServiceClient implements ExampleServiceClient.
type exampleServiceClient struct {
	exampleCall *connect.Client[test.ExampleMessage, test.ExampleMessage]
}

// ExampleCall calls test.ExampleService.ExampleCall.
func (c *exampleServiceClient) ExampleCall(ctx context.Context, req *connect.Request[test.ExampleMessage]) (*connect.Response[test.ExampleMessage], error) {
	return c.exampleCall.CallUnary(ctx, req)
}

// ExampleServiceHandler is an implementation of the test.ExampleService service.
type ExampleServiceHandler interface {
	ExampleCall(context.Context, *connect.Request[test.ExampleMessage]) (*connect.Response[test.ExampleMessage], error)
}

// NewExampleServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewExampleServiceHandler(svc ExampleServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	exampleServiceExampleCallHandler := connect.NewUnaryHandler(
		ExampleServiceExampleCallProcedure,
		svc.ExampleCall,
		connect.WithSchema(exampleServiceExampleCallMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/test.ExampleService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ExampleServiceExampleCallProcedure:
			exampleServiceExampleCallHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedExampleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedExampleServiceHandler struct{}

func (UnimplementedExampleServiceHandler) ExampleCall(context.Context, *connect.Request[test.ExampleMessage]) (*connect.Response[test.ExampleMessage], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("test.ExampleService.ExampleCall is not implemented"))
}
