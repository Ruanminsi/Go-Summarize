/**
* @Author: RMS
* @Date: 2021/3/25
 */
package main

import (
	cls "Go-Summarize/grpc/stream_client"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type SimpleService struct{}

func (s *SimpleService) Route(ctx context.Context, req *cls.SimpleRequest) (*cls.SimpleResponse, error) {
	return &cls.SimpleResponse{
		Code:  1,
		Value: "success" + req.Data,
	}, nil
}

func (s *SimpleService) RouteList(srv cls.StreamClient_RouteListServer) error {
	for {
		resp, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&cls.SimpleResponse{Code: 1, Value: "all read"})
		}
		if err != nil {
			return err
		}
		log.Println(resp.StreamData)
	}
}

const Address = "localhost:3721"

func main() {
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("err occur in listen: %v", err)
	}
	log.Println(Address + " Listening ...")
	grpcServer := grpc.NewServer()
	// grpc 服务器注册服务
	cls.RegisterStreamClientServer(grpcServer, &SimpleService{})
	// 服务器Serve（）
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("err occur in grpc.Serve: %v", err)
	}
}
