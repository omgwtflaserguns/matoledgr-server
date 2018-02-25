package transaction

import (
	"context"
	"github.com/omgwtflaserguns/matomat-server/auth"
	"github.com/omgwtflaserguns/matomat-server/db"
	pb "github.com/omgwtflaserguns/matomat-server/generated"
	"github.com/op/go-logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Service struct{}

var logger = logging.MustGetLogger("log")

func (s *Service) ListTransactions(ctx context.Context, in *pb.TransactionsRequest) (*pb.TransactionList, error) {
	login, err := auth.EnsureAuthentication(ctx)
	if err != nil {
		return &pb.TransactionList{}, status.Error(codes.Unauthenticated, "No auth cookie found, please login")
	}

	rows, err := db.DbCon.Query("SELECT t.id, t.price, t.timestamp, p.id, p.name, p.price "+
		"FROM AccountTransaction t "+
		"INNER JOIN Product p ON t.productId = p.id "+
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
	login, err := auth.EnsureAuthentication(ctx)
	if err != nil {
		return &pb.BuyResponse{}, status.Error(codes.Unauthenticated, "No auth cookie found, please login")
	}

	if in.ProductId < 1 {
		return &pb.BuyResponse{}, status.Error(codes.InvalidArgument, "invalid product-id given")
	}

	res, err := db.DbCon.Exec("INSERT INTO AccountTransaction (accountId, productId, price, timestamp) "+
		"SELECT $1, p.id, p.price, $2 "+
		"FROM product p "+
		"WHERE p.id = $3", login.User.Id, time.Now(), in.ProductId)

	rows, err2 := res.RowsAffected()

	if err != nil || err2 != nil || rows != 1 {
		logger.Errorf("Error inserting new transaction into db for user %s and product %s: %v %v", login.User.Username, in.ProductId, err, err2)
		return &pb.BuyResponse{}, status.Error(codes.Internal, "")
	}

	logger.Errorf("Error inserting new transaction into db: %v", err)
	return &pb.BuyResponse{}, nil
}
