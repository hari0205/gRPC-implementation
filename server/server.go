package main

import (
	"log"
	"net"

	"example.com/prime/prime_pb"
	"google.golang.org/grpc"
)

type Server struct {
	prime_pb.UnimplementedPrimeServiceServer
}

func (s *Server) CalculatePrime(req *prime_pb.Req, stream prime_pb.PrimeService_CalculatePrimeServer) error {
	num := req.GetNumber()
	divisor := int32(2)

	for num > 1 {
		if num%divisor == 0 {
			stream.Send(&prime_pb.Res{
				Factor: divisor,
			})
			num = num / divisor
		} else {
			divisor++
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	prime_pb.RegisterPrimeServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
