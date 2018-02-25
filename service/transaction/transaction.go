package transaction

import (
	"context"
	"github.com/omgwtflaserguns/matomat-server/auth"
	"github.com/omgwtflaserguns/matomat-server/db"
	pb "github.com/omgwtflaserguns/matomat-server/generated"
	"github.com/op/go-logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct{}

var logger = logging.MustGetLogger("log")

func (s *Service) ListTransactions(ctx context.Context, in *pb.TransactionsRequest) (*pb.TransactionList, error) {
	login, err := auth.EnsureAuthentication(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.DbCon.Query("SELECT t.id, t.price, t.timestamp, p.id, p.name, p.price "+
		"FROM AccountTransaction "+
		"WHERE accountId = $1", login.User.Id)
	defer rows.Close()

	if err != nil {
		logger.Errorf("Error at selecting transactions from db: %v", err)
		return &pb.TransactionList{}, status.Error(codes.Internal, "")
	}

	transactions := []*pb.Transaction{}
	for rows.Next() {

		t := &pb.Transaction{}
		t.Product = &pb.Product{}

		err = rows.Scan(t.Id, t.Price, t.Price, t.Product.Id, t.Product.Name, t.Product.Price)
		if err != nil {
			logger.Errorf("Error at scanning transaction from db: %v", err)
			return &pb.TransactionList{}, status.Error(codes.Internal, "")
		}
		transactions = append(transactions, t)
	}
	return &pb.TransactionList{Transactions: transactions}, nil
}

func (s *Service) Buy(ctx context.Context, in *pb.BuyRequest) (*pb.BuyResponse, error) {
	return nil, nil
}
