// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: othello/v1/othello.proto

package othelloconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	v1 "ebitengine-othello/src/gen/othello/v1"
	errors "errors"
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
	// OthelloServiceName is the fully-qualified name of the OthelloService service.
	OthelloServiceName = "othello.v1.OthelloService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OthelloServiceGetAIMoveProcedure is the fully-qualified name of the OthelloService's GetAIMove
	// RPC.
	OthelloServiceGetAIMoveProcedure = "/othello.v1.OthelloService/GetAIMove"
)

// OthelloServiceClient is a client for the othello.v1.OthelloService service.
type OthelloServiceClient interface {
	GetAIMove(context.Context, *connect.Request[v1.GetAIMoveRequest]) (*connect.Response[v1.GetAIMoveResponse], error)
}

// NewOthelloServiceClient constructs a client for the othello.v1.OthelloService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOthelloServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OthelloServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	othelloServiceMethods := v1.File_othello_v1_othello_proto.Services().ByName("OthelloService").Methods()
	return &othelloServiceClient{
		getAIMove: connect.NewClient[v1.GetAIMoveRequest, v1.GetAIMoveResponse](
			httpClient,
			baseURL+OthelloServiceGetAIMoveProcedure,
			connect.WithSchema(othelloServiceMethods.ByName("GetAIMove")),
			connect.WithClientOptions(opts...),
		),
	}
}

// othelloServiceClient implements OthelloServiceClient.
type othelloServiceClient struct {
	getAIMove *connect.Client[v1.GetAIMoveRequest, v1.GetAIMoveResponse]
}

// GetAIMove calls othello.v1.OthelloService.GetAIMove.
func (c *othelloServiceClient) GetAIMove(ctx context.Context, req *connect.Request[v1.GetAIMoveRequest]) (*connect.Response[v1.GetAIMoveResponse], error) {
	return c.getAIMove.CallUnary(ctx, req)
}

// OthelloServiceHandler is an implementation of the othello.v1.OthelloService service.
type OthelloServiceHandler interface {
	GetAIMove(context.Context, *connect.Request[v1.GetAIMoveRequest]) (*connect.Response[v1.GetAIMoveResponse], error)
}

// NewOthelloServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOthelloServiceHandler(svc OthelloServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	othelloServiceMethods := v1.File_othello_v1_othello_proto.Services().ByName("OthelloService").Methods()
	othelloServiceGetAIMoveHandler := connect.NewUnaryHandler(
		OthelloServiceGetAIMoveProcedure,
		svc.GetAIMove,
		connect.WithSchema(othelloServiceMethods.ByName("GetAIMove")),
		connect.WithHandlerOptions(opts...),
	)
	return "/othello.v1.OthelloService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OthelloServiceGetAIMoveProcedure:
			othelloServiceGetAIMoveHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOthelloServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOthelloServiceHandler struct{}

func (UnimplementedOthelloServiceHandler) GetAIMove(context.Context, *connect.Request[v1.GetAIMoveRequest]) (*connect.Response[v1.GetAIMoveResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("othello.v1.OthelloService.GetAIMove is not implemented"))
}
