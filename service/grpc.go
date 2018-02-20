package service

import (
	"log"
	"net"

	"github.com/omgwtflaserguns/matomat-server/config"
	pb "github.com/omgwtflaserguns/matomat-server/generated"
	"github.com/omgwtflaserguns/matomat-server/service/greeter"
	"github.com/omgwtflaserguns/matomat-server/service/product"
	"github.com/op/go-logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"sync"
)

var logger = logging.MustGetLogger("log")

func CreateGrpcServer(wg *sync.WaitGroup) *grpc.Server {

	conf := config.GetConfig()
	listener, err := net.Listen("tcp", conf.Grpc.Address)
	if err != nil {
		logger.Panicf("failed to listen on tcp, address: %s %v", conf.Grpc.Address, err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &greeter.Service{})
	pb.RegisterProductsServer(server, &product.Service{})

	reflection.Register(server)

	wg.Add(1)
	go runGrpcServer(server, listener, wg)

	logger.Debugf("gRPC server started at %s", conf.Grpc.Address)
	return server
}

func runGrpcServer(grpcServer *grpc.Server, listener net.Listener, wg *sync.WaitGroup) {
	defer wg.Done()

	err := grpcServer.Serve(listener)
	if err != nil {
		log.Panicf("failed to serve: %v", err)
	}
}
