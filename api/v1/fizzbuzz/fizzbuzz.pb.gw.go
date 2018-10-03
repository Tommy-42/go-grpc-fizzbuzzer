// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/v1/fizzbuzz/fizzbuzz.proto

/*
Package fizzbuzzV1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package fizzbuzzV1

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

var (
	filter_FizzBuzzService_GetFizzBuzz_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_FizzBuzzService_GetFizzBuzz_0(ctx context.Context, marshaler runtime.Marshaler, client FizzBuzzServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetFizzBuzzRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_FizzBuzzService_GetFizzBuzz_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetFizzBuzz(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_FizzBuzzService_GetFizzBuzz_1(ctx context.Context, marshaler runtime.Marshaler, client FizzBuzzServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetFizzBuzzRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["first_word"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "first_word")
	}

	protoReq.FirstWord, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "first_word", err)
	}

	val, ok = pathParams["second_word"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "second_word")
	}

	protoReq.SecondWord, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "second_word", err)
	}

	val, ok = pathParams["first_multiple"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "first_multiple")
	}

	protoReq.FirstMultiple, err = runtime.Int64(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "first_multiple", err)
	}

	val, ok = pathParams["second_multiple"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "second_multiple")
	}

	protoReq.SecondMultiple, err = runtime.Int64(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "second_multiple", err)
	}

	val, ok = pathParams["limit"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "limit")
	}

	protoReq.Limit, err = runtime.Int64(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "limit", err)
	}

	msg, err := client.GetFizzBuzz(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterFizzBuzzServiceHandlerFromEndpoint is same as RegisterFizzBuzzServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterFizzBuzzServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterFizzBuzzServiceHandler(ctx, mux, conn)
}

// RegisterFizzBuzzServiceHandler registers the http handlers for service FizzBuzzService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterFizzBuzzServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterFizzBuzzServiceHandlerClient(ctx, mux, NewFizzBuzzServiceClient(conn))
}

// RegisterFizzBuzzServiceHandler registers the http handlers for service FizzBuzzService to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "FizzBuzzServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "FizzBuzzServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "FizzBuzzServiceClient" to call the correct interceptors.
func RegisterFizzBuzzServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client FizzBuzzServiceClient) error {

	mux.Handle("GET", pattern_FizzBuzzService_GetFizzBuzz_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FizzBuzzService_GetFizzBuzz_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FizzBuzzService_GetFizzBuzz_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FizzBuzzService_GetFizzBuzz_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FizzBuzzService_GetFizzBuzz_1(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FizzBuzzService_GetFizzBuzz_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_FizzBuzzService_GetFizzBuzz_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "fizzbuzz"}, ""))

	pattern_FizzBuzzService_GetFizzBuzz_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 1, 0, 4, 1, 5, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5, 1, 0, 4, 1, 5, 6}, []string{"v1", "fizzbuzz", "first_word", "second_word", "first_multiple", "second_multiple", "limit"}, ""))
)

var (
	forward_FizzBuzzService_GetFizzBuzz_0 = runtime.ForwardResponseMessage

	forward_FizzBuzzService_GetFizzBuzz_1 = runtime.ForwardResponseMessage
)
