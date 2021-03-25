/**
* @Author: RMS
* @Date: 2021/3/25
 */
package main

import (
	hellogrpc "Go-Summarize/grpc/base"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:3721", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := hellogrpc.NewHelloWorkerClient(conn)
	resp, err := client.SayHello(context.Background(), &hellogrpc.HelloRequest{
		Say: "hello grpc server",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("grpc client result:%v\n", resp.GetReply())
}
