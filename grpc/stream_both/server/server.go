/**
* @Author: RMS
* @Date: 2021/3/25
 */
package main

import (
	bs "Go-Summarize/grpc/stream_both"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

const Address = "localhost:3721"

type StreamService struct {
}

func (s *StreamService) Route(ctx context.Context, request *bs.SimpleRequest) (*bs.SimpleResponse, error) {
	return &bs.SimpleResponse{
		Code:  1,
		Value: "success" + request.Data,
	}, nil
}

func (s *StreamService) BothRoute(srv bs.BothService_BothRouteServer) error {
	count := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&bs.BothResponse{Answer: "stream answer: " + strconv.Itoa(count) + req.Question})
		if err != nil {
			return err
		}
		count++
		log.Printf("stream client question: %s", req.Question)
	}
}

func main() {
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("err occur in listen")
	}
	log.Println(Address + " Listening...")
	grpcServer := grpc.NewServer()
	bs.RegisterBothServiceServer(grpcServer, &StreamService{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("err occur in Serve() : %v", err)
	}
}
