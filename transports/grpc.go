package transport

import (
	"context"

	"cartService/endpoints"
	pb "cartService/protos"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type gRPCServer struct {
	addItem   gt.Handler
	getCart   gt.Handler
	emptyCart gt.Handler
	*pb.UnimplementedCartServiceServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.CartServiceServer {
	return &gRPCServer{
		addItem:   gt.NewServer(endpoints.AddItem, decodeGRPCAddItemRequest, encodeGRPCAddItemResponse),
		getCart:   gt.NewServer(endpoints.GetCart, decodeGRPCGetCartRequest, encodeGRPCGetCartResponse),
		emptyCart: gt.NewServer(endpoints.EmptyCart, decodeGRPCEmptyCartRequest, encodeGRPCEmptyCartResponse),
	}
}

func (s *gRPCServer) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.Empty, error) {
	_, resp, err := s.addItem.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Empty), nil
}

func (s *gRPCServer) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.Cart, error) {
	_, resp, err := s.getCart.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Cart), nil
}

func (s *gRPCServer) EmptyCart(ctx context.Context, req *pb.EmptyCartRequest) (*pb.Empty, error) {
	_, resp, err := s.emptyCart.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Empty), nil
}

func decodeGRPCAddItemRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddItemRequest)
	return endpoints.AddItemRequest{UserId: req.UserId, Item: req.Item}, nil
}

func encodeGRPCAddItemResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &pb.Empty{}, nil
}

func decodeGRPCGetCartRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetCartRequest)
	return endpoints.GetCartRequest{UserId: req.UserId}, nil
}

func encodeGRPCGetCartResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.Cart)
	return &pb.Cart{UserId: res.UserId, Items: res.Items}, nil
}

func decodeGRPCEmptyCartRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EmptyCartRequest)
	return endpoints.EmptyCartRequest{UserId: req.UserId}, nil
}

func encodeGRPCEmptyCartResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &pb.Empty{}, nil
}
