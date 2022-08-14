package main

import (
	"context"
	"fmt"
	"io"

	"example.com/prime/prime_pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	c := prime_pb.NewPrimeServiceClient(conn)
	defer conn.Close()
	req := &prime_pb.Req{
		Number: 54546,
	}
	stream, err := c.CalculatePrime(context.Background(), req)
	if err != nil {
		panic(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("The factor: %v\n", res.GetFactor())

	}

}
