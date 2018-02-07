//go:generate protoc -I ./contracts --go_out=plugins=grpc:./generated ./contracts/matoledgr.proto

package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc/reflection"
	pb "github.com/omgwtflaserguns/matoledgr-server/generated"
	"fmt"
	"sync"
	"os"
	"time"
)

const (
	port = ":50015"
)

type server struct{}

var wg sync.WaitGroup

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Fuck off " + in.Name}, nil
}

func createGrpcServer() (*grpc.Server, net.Listener)  {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("tcp listener started")

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	return grpcServer, listener
}

func wrapGrpcServer(grpcServer *grpc.Server) {

	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithAllowedRequestHeaders([]string{"*"}))

	grpcHttpHandler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

		fmt.Printf("http: %s %s %s \n", req.Method, req.Proto, req.RequestURI)

		if wrappedGrpc.IsGrpcWebRequest(req) {
			fmt.Println("grpc detected")
			wrappedGrpc.ServeHTTP(resp, req)
		}

		// Fall back to other servers.
		http.DefaultServeMux.ServeHTTP(resp, req)
	})

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	c := cors.New(cors.Options{
		// TODO zip that
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"X-Grpc-Web", "Content-Type"},
		Debug: true,
	})

	httpServer := &http.Server{
		Addr:           ":8080",
		Handler: 		c.Handler(grpcHttpHandler),
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		IdleTimeout: 3 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog: logger,
	}

	fmt.Println("Http Server start listening...")

	wg.Add(1)
	go runHttpServer(httpServer)

	fmt.Println("Http Server started")
}

func runHttpServer(httpServer *http.Server) {
	defer wg.Done()

	err := httpServer.ListenAndServe()

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func runGrpcServer(grpcServer *grpc.Server, listener net.Listener) {
	defer wg.Done()

	err := grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	fmt.Println("Server spinning up...")

	grpcServer, listener := createGrpcServer()

	wrapGrpcServer(grpcServer)

	fmt.Println("Server wrapped")

	wg.Add(1)
	go runGrpcServer(grpcServer, listener)

	fmt.Println("Server started")
	wg.Wait()
}
