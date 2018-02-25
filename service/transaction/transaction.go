package transaction

import (
	"context"
	pb "github.com/omgwtflaserguns/matomat-server/generated"
)

type Service struct{}

func (s *Service) ListTransactions(ctx context.Context, in *pb.TransactionsRequest) (*pb.TransactionList, error) {
	return nil, nil
}

func (s *Service) Buy(ctx context.Context, in *pb.BuyRequest) (*pb.BuyResponse, error) {
	return nil, nil
}
