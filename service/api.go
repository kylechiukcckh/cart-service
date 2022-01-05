package service

import (
	"context"

	"github.com/go-kit/log"

	pb "cartService/protos"
)

type service struct {
	logger log.Logger
}

// Service interface describes a service that adds numbers
type Service interface {
	AddItem(ctx context.Context, id string, cartItem *pb.CartItem) (string, error)
	GetCart(ctx context.Context, id string) (*pb.Cart, error)
	EmptyCart(ctx context.Context, id string) (string, error)
}

// NewService returns a Service with all of the expected dependencies
func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) AddItem(ctx context.Context, id string, cartItem *pb.CartItem) (string, error) {
	return id, nil
}

func (s service) GetCart(ctx context.Context, id string) (*pb.Cart, error) {
	cartItems := []*pb.CartItem{{ProductId: "12", Quantity: 32}}
	return &pb.Cart{UserId: id, Items: cartItems}, nil
}

func (s service) EmptyCart(ctx context.Context, id string) (string, error) {
	return id, nil
}
