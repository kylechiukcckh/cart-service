package endpoints

import (
	"context"

	"cartService/service"

	pb "cartService/protos"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	AddItem   endpoint.Endpoint
	GetCart   endpoint.Endpoint
	EmptyCart endpoint.Endpoint
}

type AddItemRequest struct {
	UserId string
	Item   *pb.CartItem
}
type GetCartRequest struct {
	UserId string
}
type EmptyCartRequest struct {
	UserId string
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		AddItem:   makeAddItemEndpoint(s),
		GetCart:   makeGetCartEndpoint(s),
		EmptyCart: makeEmptyCartEndpoint(s),
	}
}

func makeAddItemEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddItemRequest)
		result, _ := s.AddItem(ctx, req.UserId, req.Item)
		return result, nil
	}
}

func makeGetCartEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetCartRequest)
		result, _ := s.GetCart(ctx, req.UserId)
		return result, nil
	}
}

func makeEmptyCartEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(EmptyCartRequest)
		result, _ := s.EmptyCart(ctx, req.UserId)
		return result, nil
	}
}
