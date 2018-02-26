package product

import (
	"context"
	"github.com/omgwtflaserguns/matomat-server/auth"
	"github.com/omgwtflaserguns/matomat-server/db"
	pb "github.com/omgwtflaserguns/matomat-server/generated"
	"github.com/op/go-logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var logger = logging.MustGetLogger("log")

type Service struct{}

func (s *Service) ListProducts(ctx context.Context, in *pb.ProductRequest) (*pb.ProductList, error) {

	_, err := auth.EnsureAuthentication(ctx)
	if err != nil {
		return &pb.ProductList{}, status.Error(codes.Unauthenticated, "No auth cookie found, please login")
	}

	rows, err := db.DbCon.Query("SELECT p.id, p.name, p.price " +
		"FROM Product p " +
		"WHERE p.isActive = 1")
	defer rows.Close()

	if err != nil {
		logger.Errorf("Error at selecting products from db: %v", err)
		return &pb.ProductList{}, status.Error(codes.Internal, "")
	}

	products := []*pb.Product{}
	for rows.Next() {

		p := &pb.Product{}

		err = rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			logger.Errorf("Error scanning products from db: %v", err)
			return &pb.ProductList{}, status.Error(codes.Internal, "")
		}
		products = append(products, p)
	}
	logger.Debugf("Returning Products: %v", products)
	return &pb.ProductList{Products: products}, nil
}
