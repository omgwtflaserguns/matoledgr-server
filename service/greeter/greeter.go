package greeter

import (
	"context"
	pb "github.com/omgwtflaserguns/matomat-server/generated"
)

type Service struct{}

func (s *Service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Fuck off " + in.Name}, nil
}
