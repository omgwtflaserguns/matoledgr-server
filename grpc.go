package main

import (
	"google.golang.org/grpc"
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	pb "github.com/omgwtflaserguns/matoledgr-server/generated"
	"context"
	"database/sql"
)

//TODO Read Port from commandline
const port = ":50015"
var dbCon *sql.DB

type grpcServer struct{}

func (s *grpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Fuck off " + in.Name}, nil
}

func (s *grpcServer) ListProducts(ctx context.Context, in *pb.ProductRequest) (*pb.ProductList, error) {
	rows, err := dbCon.Query("SELECT id, name, price FROM PRODUCT")

	if err != nil {
		panic(err)
	}

	products := []*pb.Product{}
	for rows.Next() {

		var id int32
		var name string
		var price float32

		err = rows.Scan(&id, &name, &price)
		if err != nil {
			logger.Errorf("Scan failed: %v", err)
			ctx.Err()
		}
		products = append(products, &pb.Product{Id: id, Name: name, Price: price})
	}

	return &pb.ProductList{Products: products}, nil
}

func createGrpcServer() (*grpc.Server, net.Listener)  {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Panicf("failed to listen on tcp port %s: %v", port, err)
	}
	logger.Debug("tcp listener started")

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &grpcServer{})
	pb.RegisterProductsServer(server, &grpcServer{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	return server, listener
}

func runGrpcServer(grpcServer *grpc.Server, listener net.Listener) {
	defer wg.Done()

	err := grpcServer.Serve(listener)

	if err != nil {
		log.Panicf("failed to serve: %v", err)
	}
}