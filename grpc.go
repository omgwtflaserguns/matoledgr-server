package main

import (
	"context"
	"database/sql"
	pb "github.com/omgwtflaserguns/matoledgr-server/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var dbCon *sql.DB

type grpcServer struct{}

func (s *grpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Fuck off " + in.Name}, nil
}

func (s *grpcServer) ListProducts(ctx context.Context, in *pb.ProductRequest) (*pb.ProductList, error) {
	rows, err := dbCon.Query("SELECT id, name, price FROM PRODUCT")

	if err != nil {
		logger.Panic(err)
	}

	products := []*pb.Product{}
	for rows.Next() {

		var id int32
		var name string
		var price float32

		err = rows.Scan(&id, &name, &price)
		if err != nil {
			logger.Panicf("Scan failed: %v", err)
		}
		products = append(products, &pb.Product{Id: id, Name: name, Price: price})
	}
	return &pb.ProductList{Products: products}, nil
}

func createGrpcServer() *grpc.Server {
	listener, err := net.Listen("tcp", config.grpc.address)
	if err != nil {
		logger.Panicf("failed to listen on tcp, address: %s %v", config.grpc.address, err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &grpcServer{})
	pb.RegisterProductsServer(server, &grpcServer{})

	reflection.Register(server)

	wg.Add(1)
	go runGrpcServer(server, listener)

	logger.Debugf("gRPC server started at %s", config.grpc.address)
	return server
}

func runGrpcServer(grpcServer *grpc.Server, listener net.Listener) {
	defer wg.Done()

	err := grpcServer.Serve(listener)
	if err != nil {
		log.Panicf("failed to serve: %v", err)
	}
}
