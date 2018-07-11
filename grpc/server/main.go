package main

import (
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/hxzqlh/go-experiments/grpc/customer"
)

const (
	port = ":50051"
)

// server is used to implement customer.CustomerServer.
type server struct {
	savedCustomers []*pb.Customer
}

func (s *server) CreateCustomer(ctx context.Context, c *pb.Customer) (*pb.Response, error) {
	s.savedCustomers = append(s.savedCustomers, c)
	return &pb.Response{Id: c.Id, Success: true}, nil
}

func (s *server) GetCustomers(ctx context.Context, f *pb.Filter) (*pb.Customers, error) {
	arr := []*pb.Customer{}
	for _, customer := range s.savedCustomers {
		if f.Keyword != "" {
			if !strings.Contains(customer.Name, f.Keyword) {
				continue
			}
		}
		arr = append(arr, customer)
	}
	return &pb.Customers{arr}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterGrpcserviceServer(s, &server{})
	s.Serve(lis)
}
