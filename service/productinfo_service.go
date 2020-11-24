package main

import (
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	productInfo "productinfo/service/ecommerce"
)

type server struct {
	productMap map[string]*productInfo.Product
}

func (s *server) AddProduct(ctx context.Context,
	in *productInfo.Product) (*productInfo.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating Product ID", err)
	}
	in.Id = out.String()
	log.Println(in.Id)
	if s.productMap == nil {
		s.productMap = make(map[string]*productInfo.Product)
	}
	s.productMap[in.Id] = in
	log.Println(s)
	return &productInfo.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *productInfo.ProductID) (*productInfo.Product, error) {
	value, exists := s.productMap[in.Value]
	log.Println(value)
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}