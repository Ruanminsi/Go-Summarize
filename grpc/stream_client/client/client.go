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
	"strconv"
)

var streanClient cls.StreamClientClient

const Address = "localhost:3721"

func route() {
	req := &cls.SimpleRequest{
		Data: "simple grpc",
	}
	resp, err := streanClient.Route(context.Background(), req)
	if err != nil {
		log.Fatalf("err occur in Route: %v", err)
	}
	log.Println(resp)
}

func routeList() {
	stream, err := streanClient.RouteList(context.Background())
	if err != nil {
		log.Fatalf("err occur in RouteList: %v", err)
	}
	for i := 0; i < 10; i++ {
		err = stream.Send(&cls.StreamRequest{StreamData: "streamData" + strconv.Itoa(i)})
		// 当服务端在消息没接收完前主动调用SendAndClose()关闭stream，此时客户端还执行Send()，则会返回EOF错误
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("err occur in stream request: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("err occur in RouteList get response: %v", err)
	}
	log.Println(resp)
}

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err occur in net.Conn: %v", err)
	}
	defer conn.Close()
	streanClient = cls.NewStreamClientClient(conn)
	route()
	routeList()
}
