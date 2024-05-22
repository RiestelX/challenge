package main

import (
	pb "beefcount/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type BeefServiceServer struct {
	pb.UnimplementedBeefServiceServer
}

func (s *BeefServiceServer) Summary(ctx context.Context, in *pb.Empty) (*pb.SummaryResponse, error) {
	wordCount, err := WordCount("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return nil, err
	}

	return &pb.SummaryResponse{Beef: wordCount}, nil
}

func startGRPCServer() {
	grpcServer := grpc.NewServer()
	pb.RegisterBeefServiceServer(grpcServer, &BeefServiceServer{})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("failed ", err)
		return
	}
	fmt.Println("server 50051")
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("failed ", err)
	}
}
