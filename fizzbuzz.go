package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	apiFizzBuzzV1 "github.com/tommy-42/go-grpc-fizzbuzzer/api/v1/fizzbuzz"
	svcFizzBuzzV1 "github.com/tommy-42/go-grpc-fizzbuzzer/service/v1/fizzbuzz"
)

var (
	bindGRPC = flag.String("bind_grpc", "0.0.0.0:2338", "bind address for gRPC")
	bindHTTP = flag.String("bind_http", "0.0.0.0:8080", "bind address for HTTP")

	conn *grpc.ClientConn
)

// cross-origin resource sharing
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()

		o := r.Header.Get("Origin")
		header.Set("Access-Control-Allow-Origin", o)

		if r.Method == http.MethodOptions {
			header.Set("Access-Control-Allow-Methods", strings.Join(
				[]string{
					http.MethodOptions,
					http.MethodGet,
					http.MethodPut,
					http.MethodHead,
					http.MethodPost,
					http.MethodDelete,
					http.MethodPatch,
					http.MethodTrace,
				}, ", ",
			))

			header.Set("Access-Control-Allow-Headers", strings.Join(
				[]string{
					"Access-Control-Allow-Headers",
					"Origin",
					"X-Requested-With",
					"Content-Type",
					"Accept",
				}, ", ",
			))

			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	errc := make(chan error)
	defer close(errc)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", *bindGRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fizzbuzzV1 := svcFizzBuzzV1.NewFizzBuzzService()

	grpcServer := grpc.NewServer()
	apiFizzBuzzV1.RegisterFizzBuzzServiceServer(grpcServer, fizzbuzzV1)

	go func() {
		log.Println(fmt.Sprintf("Starting GRPC server on %s", *bindGRPC))
		errc <- errors.Wrap(grpcServer.Serve(lis), "Cannot start GRPC server")
	}()

	conn, err := grpc.Dial(*bindGRPC, grpc.WithInsecure())
	if err != nil {
		panic("Couldn't contact grpc server")
	}

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))

	err = apiFizzBuzzV1.RegisterFizzBuzzServiceHandler(ctx, mux, conn)
	if err != nil {
		panic("Cannot serve fizzbuzz http api v1")
	}

	go func() {
		log.Println(fmt.Sprintf("Starting HTTP server on %s", *bindHTTP))
		handlerEnriched := cors(mux)
		errc <- errors.Wrap(http.ListenAndServe(*bindHTTP, handlerEnriched), "Cannot start HTTP server")
	}()

	fatalError := <-errc
	grpcServer.GracefulStop()
	cancel()

	if err := conn.Close(); err != nil {
		log.Println(err.Error())
	}
	if err := lis.Close(); err != nil {
		log.Println(err.Error())
	}
	log.Println(fatalError.Error())
}
