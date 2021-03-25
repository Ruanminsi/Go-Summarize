/**
* @Author: RMS
* @Date: 2021/3/25
 */
package main

import (
	server_stream "Go-Summarize/grpc/stream_server"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
)

var grpcClient server_stream.StreamServerClient

const Address = "localhost:3721"

func route() {
	req := &server_stream.SimpleRequest{
		Data: "grpc",
	}
	resp, err := grpcClient.Route(context.Background(), req)
	if err != nil {
		log.Fatalf("Call Route err: %v \n", err)
	}
	log.Println(resp)
}

func listValue() {
	req := server_stream.SimpleRequest{
		Data: "stream server grpc",
	}
	stram, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("err occur in ListValue: %v \n", err)
	}
	for {
		resp, err := stram.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListValue get stream: %v \n", err)
		}
		log.Println(resp.StreamReply)
	}
	// 可以使用CloseSend()关闭stream，这样服务端就不会继续产生流消息
	// 调用CloseSend()后，若继续调用Recv()，会重新激活stream，接着之前结果获取消息
	// stream.CloseSend()
}

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err occur in net.Conn: %v \n", err)
	}
	defer conn.Close()
	grpcClient = server_stream.NewStreamServerClient(conn)
	route()
	listValue()
}
