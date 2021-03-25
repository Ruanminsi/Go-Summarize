/**
* @Author: RMS
* @Date: 2021/3/25
 */
package main

import (
	hellogrpc "Go-Summarize/grpc/base"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	Addr = "localhost:3721"
)

type HelloServiceImpl struct {
}

func (p *HelloServiceImpl) SayHello(ctx context.Context, req *hellogrpc.HelloRequest) (
	resp *hellogrpc.HelloResponse, err error) {
	resp = new(hellogrpc.HelloResponse)
	resp.Reply = "hello grpc client."
	log.Println("req: ", req.Say)
	return resp, nil
}
func main() {
	grpcServer := grpc.NewServer()
	hellogrpc.RegisterHelloWorkerServer(grpcServer, new(HelloServiceImpl))
	lst, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		fmt.Printf("服务端的gRPC进程服务 %d %s", os.Getpid(), Addr)
	}()
	log.Fatal(grpcServer.Serve(lst))
}
