/**
* @Author: RMS
* @Date: 2021/3/25
 */
package main

import (
	server_stream "Go-Summarize/grpc/stream_server"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const Addr = "localhost:3721"

type ServerStream struct {
}

func (s *ServerStream) Route(ctx context.Context, req *server_stream.SimpleRequest) (*server_stream.SimpleResponse, error) {
	return &server_stream.SimpleResponse{Code: "200", Val: "hello " + req.Data}, nil
}

func (s *ServerStream) ListValue(req *server_stream.SimpleRequest, srv server_stream.StreamServer_ListValueServer) error {
	for i := 0; i < 10; i++ {
		err := srv.Send(&server_stream.StreamResponse{
			StreamReply: req.Data + strconv.Itoa(i),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalf("err occur in listen: %v", err)
	}
	grcpServer := grpc.NewServer()
	server_stream.RegisterStreamServerServer(grcpServer, &ServerStream{})
	err = grcpServer.Serve(listener)
	if err != nil {
		log.Fatalf(" err occur in grpcServer.Serve : %v", err)
	}

}
