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
	"strconv"
)

const Address = "localhost:3721"

var streamClient bs.BothServiceClient

func route() {
	req := &bs.SimpleRequest{
		Data: "simple grpc",
	}
	resp, err := streamClient.Route(context.Background(), req)
	if err != nil {
		log.Fatalf("err occur in route: %v", err)
	}
	log.Println(resp)
}

func bothRoute() {
	stream, err := streamClient.BothRoute(context.Background())
	if err != nil {
		log.Fatalf("err occur in bothRoute get stream: %v", err)
	}
	for i := 0; i < 10; i++ {
		err = stream.Send(&bs.BothRequest{Question: "client send stream " + strconv.Itoa(i)})
		if err != nil {
			log.Fatalf("err occur in request: %v", err)
		}
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("err occur in get: %v", err)
		}
		log.Println(resp)
	}
}

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err occur in net.Conn: %v", err)
	}
	defer conn.Close()
	streamClient = bs.NewBothServiceClient(conn)
	route()
	bothRoute()
}
